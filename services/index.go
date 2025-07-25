package services

import (
	services_redis "github.com/ballinwza/combine-be-workshop/services/redis"
	"github.com/redis/go-redis/v9"
)

type InjectorServices struct {
	RedisService *services_redis.RedisService
}

func NewInjectorServices() *InjectorServices {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	redis := services_redis.NewRedisService(rdb)

	return &InjectorServices{
		RedisService: redis,
	}
}
