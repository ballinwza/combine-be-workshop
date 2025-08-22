package services_redis_ticket

import (
	"github.com/redis/go-redis/v9"
)

type RedisTicketService struct {
	rdb           *redis.Client
	ticketKey     string
	ticketChannel string
}

func NewRedisTicketService(rdb *redis.Client) *RedisTicketService {
	ticketKey := "available_ticket"
	ticketChannel := "ticket_remaining_channel"

	return &RedisTicketService{
		rdb:           rdb,
		ticketKey:     ticketKey,
		ticketChannel: ticketChannel,
	}
}
