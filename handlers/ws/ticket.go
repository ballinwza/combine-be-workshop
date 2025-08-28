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

		// Initial
		initialData, err := h.RedisTicketService.SaveRemainingTicket(ctx)
		if err != nil {
			log.Printf("Error WsTicket First fetch : %v\n", err)
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

		// Notification
		bookingTicketList, err := h.RedisTicketService.BookingTicketList(ctx, userId)
		if err != nil {
			return
		}

		jsonData, _ := json.Marshal(bookingTicketList)
		err = c.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			log.Printf("Failed to send pending message to user %s: %v", userId, err)
			return
		}

		notification, _ := h.RedisTicketService.GetNotification(ctx, userId)
		if notification != nil {
			jsonData, _ := json.Marshal(notification)
			err := c.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				log.Printf("Failed to send pending message to user %s: %v", userId, err)
				return
			}

			err = h.RedisTicketService.DeleteNotification(ctx, userId)
			if err != nil {
				log.Printf("Failed to send pending message to user %s: %v", userId, err)
				return
			}

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
