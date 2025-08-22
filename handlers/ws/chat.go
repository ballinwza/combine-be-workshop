package handlers_ws

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ballinwza/combine-be-workshop/services"
	services_redis_chat "github.com/ballinwza/combine-be-workshop/services/redis/chat"
	"github.com/gofiber/contrib/websocket"
)

type WsChatHandler struct {
	ChatService *services_redis_chat.RedisChatService
}

func NewWsChatHandler() WsChatHandler {
	redisChatService := services.NewInjectorServices(nil).RedisServices.RedisChatServices
	return WsChatHandler{
		ChatService: redisChatService,
	}
}

func (h *WsChatHandler) WsChat() func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		ctx := context.Background()

		go func() {
			initChat, err := h.ChatService.GetChatHistory(ctx)
			if err == nil {
				jsonData, _ := json.Marshal(initChat)
				if err := c.WriteMessage(websocket.TextMessage, jsonData); err != nil {
					log.Printf("ส่งข้อมูลชุดแรกไม่สำเร็จ: %v", err)
				}
			}

			updateChane, err := h.ChatService.SubscribeChat(ctx)
			if err != nil {
				log.Printf("ไม่สามารถ Subscribe ได้: %v", err)
				c.Close()
				return
			}

			for newMessage := range updateChane {
				jsonData, _ := json.Marshal(newMessage)
				if err := c.WriteMessage(websocket.TextMessage, jsonData); err != nil {
					log.Printf("ส่งข้อความหา Client ไม่สำเร็จ: %v", err)
					break
				}

			}
		}()

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Printf("Client ปิดการเชื่อมต่อ: %s", c.RemoteAddr())
				break
			}

			if err := h.ChatService.PublishChat(ctx, msg); err != nil {
				log.Printf("Publish ไปยัง Redis ไม่สำเร็จ: %v", err)
				continue
			}

		}
	}
}
