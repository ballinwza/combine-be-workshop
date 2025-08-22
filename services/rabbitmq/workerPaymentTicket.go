package services_rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	services_redis_ticket "github.com/ballinwza/combine-be-workshop/services/redis/ticket"
)

func (s *RabbitmqService) WorkerPaymentTicket() (<-chan struct{}, error) {

	err := s.exchangeBookingDeclare()
	if err != nil {
		return nil, err
	}

	q, err := s.ch.QueueDeclare(
		"ticket_payment",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = s.ch.QueueBind(
		q.Name,
		"",
		"ticket",
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = s.ch.Qos(1, 0, false)
	if err != nil {
		return nil, err
	}

	msgs, err := s.ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			userId := string(d.Body)

			// IsPending
			initialTicket, err := s.redis.RedisTicketServices.CreateTicket(context.Background())
			if err != nil {
				fmt.Printf("Error")
			}

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

			ticketPayload := services_redis_ticket.RemainingTicket{
				RemainingTicket: initialTicket.RemainingTicket,
				Type:            "ticket_remaining",
			}
			ticketPayloadByte, err := json.Marshal(ticketPayload)
			if err != nil {
				fmt.Printf("Error")
			}

			s.pool.SendToUser(userId, payload)
			s.pool.Broadcast(ticketPayloadByte)
			// TODO: ต้องเปลี่ยนไปใช้ Mutex sendToUser แทน
			// err = s.redis.RedisTicketServices.PublishRemainingTicket(context.Background(), bytePayload)
			// if err != nil {
			// 	fmt.Printf("Error WorkerPaymentTicket PublishRemainingTicket : %v\n", err)
			// }

			time.Sleep(5 * time.Second)
			remainingTickets, err := s.redis.RedisTicketServices.DescreaseTicket(context.Background())
			if err != nil {
				fmt.Printf("Error WorkerPaymentTicket DescreaseTicket : %s\n", err)
				d.Nack(false, true)
				continue
			}
			fmt.Printf("Done\n")
			fmt.Printf("Success WorkerPaymentTicket remainingTickets is : %v\n", *remainingTickets)
			d.Ack(false)

			// payload := fmt.Sprintf(`{"leftTicket": %d}`, *remainingTickets)

			// Successful
			successBookingStatusPayload := services_redis_ticket.NotificationTicket{
				UserId:    userId,
				Status:    true,
				IsPending: false,
				Message:   "Your ticket booking was successful!",
				Type:      "booking_status",
			}
			jsonData, err := json.Marshal(successBookingStatusPayload)
			if err != nil {
				fmt.Printf("Error WorkerPaymentTicket marshal data : %v\n", err)
			}

			jsonTicketLeft, err := json.Marshal(services_redis_ticket.RemainingTicket{
				RemainingTicket: *remainingTickets,
				Type:            "ticket_remaining",
			})
			if err != nil {
				fmt.Printf("Error WorkerPaymentTicket marshal data : %v\n", err)
			}

			s.pool.SendToUser(userId, jsonData)
			s.pool.Broadcast(jsonTicketLeft)
			// TODO: ต้องเปลี่ยนไปใช้ Mutex sendToUser แทน
			// err = s.redis.RedisTicketServices.PublishRemainingTicket(context.Background(), jsonData)
			// if err != nil {
			// 	fmt.Printf("Error WorkerPaymentTicket PublishRemainingTicket : %v\n", err)
			// }
		}
	}()

	return forever, nil
}
