package services_redis_ticket

import (
	"github.com/redis/go-redis/v9"
)

type RedisTicketService struct {
	rdb *redis.Client
}

func NewRedisTicketService(rdb *redis.Client) *RedisTicketService {

	return &RedisTicketService{
		rdb: rdb,
	}
}
