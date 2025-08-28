package services_redis_ticket

import (
	"context"
	"fmt"
)

func (s *RedisTicketService) DeleteNotification(ctx context.Context, userId string) error {
	redisKey := "notification:" + userId

	err := s.rdb.Del(ctx, redisKey).Err()
	if err != nil {
		fmt.Printf("Error get notification \n")
		return err
	}

	return nil
}
