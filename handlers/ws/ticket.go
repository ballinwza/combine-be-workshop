package handlers_ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	handlers_mutex "github.com/ballinwza/combine-be-workshop/handlers/mutex"
	"github.com/ballinwza/combine-be-workshop/services"

	services_redis_ticket "github.com/ballinwza/combine-be-workshop/services/redis/ticket"
	"github.com/gofiber/contrib/websocket"
)

type WsTicketHandler struct {
	RedisTicketService *services_redis_ticket.RedisTicketService
	pool               *handlers_mutex.ClientPool
}

func NewWsTicketHandler(pool *handlers_mutex.ClientPool) WsTicketHandler {
	redisTicketService := services.NewInjectorServices(pool).RedisServices.RedisTicketServices

	return WsTicketHandler{
		RedisTicketService: redisTicketService,
		pool:               pool,
	}
}

func (h *WsTicketHandler) WsTicket() func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		userId := c.Params("userId")

		h.pool.Add(userId, c)
		defer h.pool.Remove(userId)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var resultData services_redis_ticket.NotificationTicket

		// Initial
		initialData, err := h.RedisTicketService.CreateTicket(ctx)
		if err != nil {
			log.Printf("Error WsTicket First fetch : %v\n", err)
		}

		resultData = services_redis_ticket.NotificationTicket{
			UserId:    userId,
			Status:    false,
			IsPending: false,
			Message:   "Not on queue",
			Type:      "booking_status",
		}

		jsonData, _ := json.Marshal(resultData)
		if err := c.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			fmt.Printf("Error WsTicket sending data : %v\n", err)
			c.Close()
			return
		}

		ticketPayload := services_redis_ticket.RemainingTicket{
			RemainingTicket: initialData.RemainingTicket,
			Type:            "ticket_remaining",
		}
		ticketPayloadByte, err := json.Marshal(ticketPayload)
		if err != nil {
			fmt.Printf("Error")
		}

		err = c.WriteMessage(websocket.TextMessage, ticketPayloadByte)
		if err != nil {
			fmt.Printf("Error WsTicket sending data : %v\n", err)
			c.Close()
			return
		}

		// Waiting
		for {
			if _, _, err := c.ReadMessage(); err != nil {
				// Error นี้จะเกิดขึ้นเมื่อ Client ปิด Browser หรือเน็ตหลุด
				log.Printf("User '%s' connection closed.", userId)
				break // ออกจาก Loop แล้ว defer จะทำงาน
			}
		}

	}
}
