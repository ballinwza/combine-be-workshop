package handlers_ws

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ballinwza/combine-be-workshop/services"
	services_redis_leaderboard "github.com/ballinwza/combine-be-workshop/services/redis/leaderboard"
	"github.com/gofiber/contrib/websocket"
)

type WsLeaderboardHandler struct {
	RedisLeaderboardService *services_redis_leaderboard.RedisLeaderboardService
}

func NewWsLeaderboardHandler(allService services.InjectorServices) WsLeaderboardHandler {
	redisLeaderboardService := allService.RedisServices.RedisLeaderboardServices

	return WsLeaderboardHandler{
		RedisLeaderboardService: redisLeaderboardService,
	}
}

func (h *WsLeaderboardHandler) WsLeaderboardScore() func(c *websocket.Conn) {

	return func(c *websocket.Conn) {
		ctx := context.Background()

		initialData, err := h.RedisLeaderboardService.GetSortScore(ctx)
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

		updateChan, err := h.RedisLeaderboardService.SubscribeLeaderboard(ctx)
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
