package services_redis_ticket

import (
	"context"
	"encoding/json"
)

func (s *RedisTicketService) SaveNotification(ctx context.Context, userId string, isSuccess bool, message string) (*NotificationTicket, error) {
	redisKey := "notification:" + userId

	notificationPayload := NotificationTicket{
		IsSuccess: isSuccess,
		Message:   message,
		Type:      Notification,
	}

	jsonData, _ := json.Marshal(notificationPayload)
	err := s.rdb.Set(ctx, redisKey, string(jsonData), -1).Err()
	if err != nil {
		return nil, err
	}

	return &notificationPayload, err
}
