package services_redis_ticket

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *RedisTicketService) GetNotification(ctx context.Context, userId string) (*NotificationTicket, error) {
	redisKey := "notification:" + userId
	var payload NotificationTicket

	message, err := s.rdb.Get(ctx, redisKey).Result()
	if err != nil {
		fmt.Printf("Error GetNotification GET : %v\n", err)
		return nil, err
	}

	err = json.Unmarshal([]byte(message), &payload)
	if err != nil {
		fmt.Printf("Error GetNotification Unmarshal : %v\n", err)
		return nil, err
	}
	return &payload, nil
}
