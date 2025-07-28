package handlers_ws

import (
	"context"
	"encoding/json"
	"log"

	services_redis_leaderboard "github.com/ballinwza/combine-be-workshop/services/redis/leaderboard"
	"github.com/gofiber/contrib/websocket"
)

func WsLeaderboardScore(redisLeaderboardService *services_redis_leaderboard.RedisLeaderboardService) func(c *websocket.Conn) {

	return func(c *websocket.Conn) {
		ctx := context.Background()

		initialData, err := redisLeaderboardService.GetSortScore(ctx)
		if err != nil {
			log.Printf("ดึงข้อมูลครั้งแรกไม่สำเร็จ: %v", err)
		} else {
			jsonData, _ := json.Marshal(initialData)
			if err := c.WriteMessage(websocket.TextMessage, jsonData); err != nil {
				log.Printf("ส่งข้อมูลครั้งแรกไม่สำเร็จ: %v", err)
				c.Close()
				return
			}
		}

		updateChan, err := redisLeaderboardService.SubscribeLeaderboard(ctx)
		if err != nil {
			log.Printf("ไม่สามารถ Subscribe ได้: %v", err)
			c.Close()
			return
		}

		for newLeaderboard := range updateChan {
			jsonData, _ := json.Marshal(newLeaderboard)
			if err := c.WriteMessage(websocket.TextMessage, jsonData); err != nil {
				log.Printf("ส่งข้อมูลหา Client ไม่ได้แล้ว: %v", err)
				break
			}
		}

	}
}
