package services_redis

import (
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	rdb *redis.Client
}

func NewRedisService(rdb *redis.Client) *RedisService {
	return &RedisService{
		rdb: rdb,
	}
}
