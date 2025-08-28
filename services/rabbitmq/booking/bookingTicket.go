package services_rabbitmq_booking

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	services_redis_ticket "github.com/ballinwza/combine-be-workshop/services/redis/ticket"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rabbitmq/amqp091-go"
)

func (s *RabbitMqBookingService) BookingTicket(ctx context.Context, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	generateNewId, _ := gonanoid.Generate("0123456789", 10)
	bookingPayload, err := s.redis.RedisTicketServices.SaveBookingTicket(ctx, generateNewId, userId, false, true, "กำลังดำเนินการจองตั๋ว", services_redis_ticket.Pending)
	if err != nil {
		fmt.Printf("Error : %v/", err)
	}
	payload, err := json.Marshal(bookingPayload)
	if err != nil {
		fmt.Printf("Error")
	}
	s.pool.SendToUser(userId, payload)

	err = s.ch.ExchangeDeclare(
		BOOKING_EXCHANGE,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Error exchangeBookingDeclare : %s\n", err)
		return err
	}

	err = s.ch.PublishWithContext(ctx,
		BOOKING_EXCHANGE,
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
