package handlers_ws

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/redis/go-redis/v9"
)

const redisChannel = "chat_messages"

func TestPubsub(c *websocket.Conn) {

	var (
		mt  int
		msg []byte
		// err error
	)
	log.Printf("✅ Client MSG: %v", msg)
	log.Printf("✅ Client MT: %v", mt)

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	go func() {
		pubsub := rdb.Subscribe(ctx, redisChannel)
		_, err := pubsub.Receive(ctx)
		if err != nil {
			log.Printf("เกิดข้อผิดพลาดตอน Subscribe: %v", err)
			return
		}

		ch := pubsub.Channel()

		for msg := range ch {
			if err := c.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
				log.Printf("ส่งข้อความหา Client ไม่สำเร็จ: %v", err)
				break
			}

			fmt.Printf("Log %v", msg.Payload)
		}
	}()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			// if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
			// 	log.Printf("Client ปิดการเชื่อมต่อ: %s", c.RemoteAddr())
			// }
			log.Printf("Client ปิดการเชื่อมต่อ: %s", c.RemoteAddr())
			break
		}

		if err := rdb.Publish(ctx, redisChannel, msg).Err(); err != nil {
			log.Printf("Publish ไปยัง Redis ไม่สำเร็จ: %v", err)
			continue
		}
	}
}
