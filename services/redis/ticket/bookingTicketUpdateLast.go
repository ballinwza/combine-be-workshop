package services_redis_ticket

import (
	"context"
	"encoding/json"
)

func (s *RedisTicketService) UpdateLastBookingTicket(ctx context.Context, userId string, message string) (*BookingTicket, error) {
	redisKey := "booking_ticket:" + userId
	var payload BookingTicket

	lastIndexValue, err := s.rdb.RPop(ctx, redisKey).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(lastIndexValue), &payload)
	if err != nil {
		return nil, err
	}

	result, err := s.SaveBookingTicket(ctx, *payload.Id, userId, true, false, message, Successed)
	if err != nil {
		return nil, err
	}

	return result, nil
}
