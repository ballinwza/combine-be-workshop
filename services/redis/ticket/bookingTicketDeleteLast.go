package services_redis_ticket

import (
	"context"
)

func (s *RedisTicketService) DeleteLastBookingTicket(ctx context.Context, userId string) error {
	redisKey := "booking_ticket:" + userId

	err := s.rdb.RPop(ctx, redisKey).Err()
	if err != nil {
		return err
	}

	return nil

}
