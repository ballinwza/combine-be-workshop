package services_redis_ticket

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *RedisTicketService) SaveNotification(ctx context.Context, userId string, isSuccess bool, message string) (*NotificationTicket, error) {
	redisKey := "notification:" + userId

	notificationPayload := NotificationTicket{
		IsSuccess: isSuccess,
		Message:   message,
		Type:      Notification,
	}

	jsonData, err := json.Marshal(notificationPayload)
	if err != nil {
		fmt.Printf("Error SaveNotification Marshal : %v\n", err)
		return nil, err
	}

	err = s.rdb.Set(ctx, redisKey, string(jsonData), -1).Err()
	if err != nil {
		fmt.Printf("Error SaveNotification UPDATE : %v\n", err)
		return nil, err
	}

	return &notificationPayload, err
}
