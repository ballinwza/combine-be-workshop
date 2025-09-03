package services_redis_chat

import (
	"github.com/redis/go-redis/v9"
)

type RedisChatService struct {
	rdb *redis.Client
}

func NewRedisChatService(rdb *redis.Client) *RedisChatService {

	return &RedisChatService{
		rdb: rdb,
	}
}
