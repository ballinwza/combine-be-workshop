package services_redis_ticket

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *RedisTicketService) BookingTicketList(ctx context.Context, userId string) ([]BookingTicket, error) {
	redisKey := "booking_ticket:" + userId
	var result []BookingTicket

	messages, err := s.rdb.LRange(ctx, redisKey, 0, -1).Result()
	if err != nil {
		fmt.Printf("Error BookingTicketList GET : %v\n", err)
		return nil, err
	}

	for _, value := range messages {
		var temporaryNotiData BookingTicket
		err := json.Unmarshal([]byte(value), &temporaryNotiData)
		if err != nil {
			fmt.Printf("Error BookingTicketList Unmarshal : %v\n", err)
			return nil, err
		}

		result = append(result, temporaryNotiData)
	}

	return result, nil
}
