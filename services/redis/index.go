package services_redis

import (
	"os"

	services_redis_cache "github.com/ballinwza/combine-be-workshop/services/redis/cache"
	services_redis_chat "github.com/ballinwza/combine-be-workshop/services/redis/chat"
	services_redis_leaderboard "github.com/ballinwza/combine-be-workshop/services/redis/leaderboard"
	services_redis_ticket "github.com/ballinwza/combine-be-workshop/services/redis/ticket"
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	RedisLeaderboardServices *services_redis_leaderboard.RedisLeaderboardService
	RedisChatServices        *services_redis_chat.RedisChatService
	RedisTicketServices      *services_redis_ticket.RedisTicketService
	RedisCacheServices       *services_redis_cache.RedisCache
}

func NewRedisService() *RedisService {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	redisLeaderboard := services_redis_leaderboard.NewRedisLeaderboardService(rdb)
	redisChat := services_redis_chat.NewRedisChatService(rdb)
	redisTicket := services_redis_ticket.NewRedisTicketService(rdb)
	redisCache := services_redis_cache.NewRedisCache(rdb)

	return &RedisService{
		RedisLeaderboardServices: redisLeaderboard,
		RedisChatServices:        redisChat,
		RedisTicketServices:      redisTicket,
		RedisCacheServices:       redisCache,
	}
}
