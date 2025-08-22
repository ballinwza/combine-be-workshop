package services_rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	services_redis_ticket "github.com/ballinwza/combine-be-workshop/services/redis/ticket"
	"github.com/rabbitmq/amqp091-go"
)

func (s *RabbitmqService) BookingTicket(ctx context.Context, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	bookingPayload := services_redis_ticket.NotificationTicket{
		UserId:    userId,
		Status:    true,
		IsPending: true,
		Message:   "Your ticket booking is pending",
		Type:      "booking_status",
	}
	payload, err := json.Marshal(bookingPayload)
	if err != nil {
		fmt.Printf("Error")
	}
	s.pool.SendToUser(userId, payload)

	err = s.exchangeBookingDeclare()
	if err != nil {
		return err
	}

	// fmt.Printf("Publishing Key : %s\n", event.RoutingKey)

	err = s.ch.PublishWithContext(ctx,
		"ticket",
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(userId),
		})
	if err != nil {
		fmt.Printf("Failed to open channel\n")
		return err
	}

	return nil
}
