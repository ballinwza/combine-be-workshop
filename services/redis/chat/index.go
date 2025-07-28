package services_redis_chat

import (
	"github.com/redis/go-redis/v9"
)

type RedisChatService struct {
	rdb            *redis.Client
	channel        string
	chatHistoryKey string
}

func NewRedisChatService(rdb *redis.Client) *RedisChatService {
	chatChennel := "chat_channel"
	chatKey := "chat_history"

	return &RedisChatService{
		rdb:            rdb,
		channel:        chatChennel,
		chatHistoryKey: chatKey,
	}
}
