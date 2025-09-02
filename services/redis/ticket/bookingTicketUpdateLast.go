package services_redis_ticket

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *RedisTicketService) UpdateLastBookingTicket(ctx context.Context, userId string, message string) (*BookingTicket, error) {
	redisKey := "booking_ticket:" + userId
	var payload BookingTicket

	lastIndexValue, err := s.rdb.RPop(ctx, redisKey).Result()
	if err != nil {
		fmt.Printf("Error UpdateLastBookingTicket pop list : %v\n", err)
		return nil, err
	}

	err = json.Unmarshal([]byte(lastIndexValue), &payload)
	if err != nil {
		fmt.Printf("Error UpdateLastBookingTicket Unmarshal : %v\n", err)
		return nil, err
	}

	result, err := s.SaveBookingTicket(ctx, *payload.Id, userId, true, false, message, Successed)
	if err != nil {
		fmt.Printf("Error UpdateLastBookingTicket UPDATE : %v\n", err)
		return nil, err
	}

	return result, nil
}
