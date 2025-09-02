package services_redis_ticket

import (
	"context"
	"fmt"
)

func (s *RedisTicketService) DeleteLastBookingTicket(ctx context.Context, userId string) error {
	redisKey := "booking_ticket:" + userId

	err := s.rdb.RPop(ctx, redisKey).Err()
	if err != nil {
		fmt.Printf("Error DeleteLastBookingTicket : %v\n", err)
		return err
	}

	return nil

}
