package services_rabbitmq_booking

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	services_redis_ticket "github.com/ballinwza/combine-be-workshop/services/redis/ticket"
	"github.com/rabbitmq/amqp091-go"
)

func (s *RabbitMqBookingService) WorkerPaymentTicket() (<-chan struct{}, error) {
	// Queue Config
	err := s.ch.Qos(1, 0, false)
	if err != nil {
		return nil, err
	}

	msgs, err := s.ch.Consume(
		BOOKING_QUEUE,
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

	// Start Rouetine Process
	var forever chan struct{}

	go func() {
		var remainingPayload services_redis_ticket.RemainingTicket

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for d := range msgs {
			userId := string(d.Body)

			// IsPending
			initialTicket, err := s.redis.RedisTicketServices.SaveRemainingTicket(context.Background())
			if err != nil {
				fmt.Printf("Error")
			}
			remainingPayload = services_redis_ticket.RemainingTicket{
				RemainingTicket: initialTicket.RemainingTicket,
				Type:            services_redis_ticket.Remaining,
			}

			ticketPayloadByte, err := json.Marshal(remainingPayload)
			if err != nil {
				fmt.Printf("Error")
			}

			s.pool.Broadcast(ticketPayloadByte)
			time.Sleep(5 * time.Second)

			// Task working
			if initialTicket.RemainingTicket <= 0 {

				err := s.redis.RedisTicketServices.DeleteLastBookingTicket(ctx, userId)
				if err != nil {
					fmt.Printf("Error : %v/n", err)
				}

				bookingticketList, err := s.redis.RedisTicketServices.BookingTicketList(ctx, userId)
				if err != nil {
					fmt.Printf("Error : %v/n", err)
				}
				jsonData, _ := json.Marshal(bookingticketList)
				s.pool.SendToUser(userId, jsonData)

				notiSuccess, _ := s.redis.RedisTicketServices.SaveNotification(ctx, userId, false, "ตั๋วหมดไอหนุ่ม")
				jsonNotiSuccess, _ := json.Marshal(notiSuccess)
				s.pool.SendToUser(userId, jsonNotiSuccess)
				s.redis.RedisTicketServices.DeleteNotification(ctx, userId)

				jsonTicketLeft, err := json.Marshal(remainingPayload)
				if err != nil {
					fmt.Printf("Error WorkerPaymentTicket marshal data : %v\n", err)
				}
				s.pool.Broadcast(jsonTicketLeft)

				d.Ack(false)
				continue
			}

			remainingTickets, err := s.redis.RedisTicketServices.DescreaseTicket(context.Background())
			// Successful
			if err == nil {
				fmt.Printf("User Succes : %v\n", userId)
				payloadAfterSave, _ := s.redis.RedisTicketServices.UpdateLastBookingTicket(ctx, userId, "จองตั๋วสำเร็จ")
				jsonData, err := json.Marshal(payloadAfterSave)
				if err != nil {
					fmt.Printf("Error WorkerPaymentTicket marshal data : %v\n", err)
				}

				fmt.Printf("Done\n")
				fmt.Printf("Success WorkerPaymentTicket remainingTickets is : %v\n", *remainingTickets)
				d.Ack(false)

				notiSuccess, _ := s.redis.RedisTicketServices.SaveNotification(ctx, userId, true, "จองตั๋วสำเร็จแล้ว")
				jsonNotiSuccess, _ := json.Marshal(notiSuccess)
				s.pool.SendToUser(userId, jsonNotiSuccess)

				s.pool.SendToUser(userId, jsonData)

				jsonTicketLeft, err := json.Marshal(services_redis_ticket.RemainingTicket{
					RemainingTicket: *remainingTickets,
					Type:            services_redis_ticket.Remaining,
				})
				if err != nil {
					fmt.Printf("Error WorkerPaymentTicket marshal data : %v\n", err)
				}
				s.pool.Broadcast(jsonTicketLeft)
				continue
			}

			// Handle error and retry section
			retryCount, _ := d.Headers["x-retry-count"].(int32)

			fmt.Printf("Task for retry %v \n", retryCount)

			if retryCount >= MAX_RETRIES {
				fmt.Printf("Error WorkerPaymentTicket Permanant delete : %s\n", err)
				d.Nack(false, false)
			} else {
				fmt.Printf("Error WorkerPaymentTicket with Retry : %s\n", err)

				newHeaders := amqp091.Table{
					"x-retry-count": retryCount + 1,
				}

				err := s.ch.PublishWithContext(
					context.Background(),
					RETRY_BOOKING_EXCHANGE,
					RETRY_BOOKING_QUEUE,
					false,
					false,
					amqp091.Publishing{
						ContentType: d.ContentType,
						Body:        d.Body,
						Headers:     newHeaders,
					},
				)
				if err != nil {
					fmt.Printf("FATAL: Failed to publish message to retry exchange: %v. Message sent to DLQ.\n", err)
					d.Nack(false, false) // ถ้า Publish ไป Retry ไม่ได้ ก็ส่งไป DLQ แทน
					continue
				} else {
					// **สำคัญมาก:** เมื่อส่งไป Retry สำเร็จ ให้ Ack message เก่าทิ้ง
					// เพราะเราได้สร้าง message ใหม่ในระบบ Retry แล้ว
					notiSuccess, _ := s.redis.RedisTicketServices.SaveNotification(ctx, userId, false, "ไม่สามารถจองตั๋วได้ไม่รู้เป็นอะไร")
					jsonNotiSuccess, _ := json.Marshal(notiSuccess)
					s.pool.SendToUser(userId, jsonNotiSuccess)
					s.redis.RedisTicketServices.DeleteNotification(ctx, userId)
					d.Ack(false)
					continue
				}
			}

		}
	}()

	return forever, nil
}
