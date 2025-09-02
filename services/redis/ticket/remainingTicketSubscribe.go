package services_redis_ticket

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *RedisTicketService) SubscribeRemainingTicket(ctx context.Context) (<-chan NotificationTicket, error) {
	channel := make(chan NotificationTicket)

	sub := s.rdb.Subscribe(ctx, ticketChannel)
	_, err := sub.Receive(ctx)
	if err != nil {
		fmt.Printf("Error SubscribeRemainingTicket : %v\n", err)
		return nil, err
	}

	go func() {
		defer close(channel)
		defer sub.Close()

		ch := sub.Channel()

		for msg := range ch {
			var ticket NotificationTicket
			err := json.Unmarshal([]byte(msg.Payload), &ticket)
			if err != nil {
				fmt.Printf("Error SubscribeRemainingTicket : %v\n", err)
				continue
			}

			channel <- ticket

		}
	}()

	return channel, nil
}
