package services

import (
	handlers_mutex "github.com/ballinwza/combine-be-workshop/handlers/mutex"
	services_rabbitmq "github.com/ballinwza/combine-be-workshop/services/rabbitmq"
	services_redis "github.com/ballinwza/combine-be-workshop/services/redis"
	"github.com/redis/go-redis/v9"
)

type InjectorServices struct {
	RedisServices  *services_redis.RedisService
	RabbitServices *services_rabbitmq.RabbitmqService
}

func NewInjectorServices(pool *handlers_mutex.ClientPool, rdb *redis.Client) *InjectorServices {
	redisService := services_redis.NewRedisService(rdb)
	rabbitServices := services_rabbitmq.NewRabbitmqService(pool, redisService)

	return &InjectorServices{
		RedisServices:  redisService,
		RabbitServices: rabbitServices,
	}
}
