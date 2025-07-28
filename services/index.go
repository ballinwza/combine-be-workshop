package services

import (
	services_redis "github.com/ballinwza/combine-be-workshop/services/redis"
	services_redis_leaderboard "github.com/ballinwza/combine-be-workshop/services/redis/leaderboard"
	"github.com/redis/go-redis/v9"
)

type InjectorServices struct {
	RedisService             *services_redis.RedisService
	RedisLeaderboardServices *services_redis_leaderboard.RedisLeaderboardService
}

func NewInjectorServices() *InjectorServices {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	redis := services_redis.NewRedisService(rdb)
	redisLeaderboard := services_redis_leaderboard.NewRedisLeaderboardService(rdb)

	return &InjectorServices{
		RedisService:             redis,
		RedisLeaderboardServices: redisLeaderboard,
	}
}
