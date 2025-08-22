package services

import (
	handlers_mutex "github.com/ballinwza/combine-be-workshop/handlers/mutex"
	services_rabbitmq "github.com/ballinwza/combine-be-workshop/services/rabbitmq"
	services_redis "github.com/ballinwza/combine-be-workshop/services/redis"
)

type InjectorServices struct {
	RedisServices  *services_redis.RedisService
	RabbitServices *services_rabbitmq.RabbitmqService
}

func NewInjectorServices(pool *handlers_mutex.ClientPool) *InjectorServices {
	redisService := services_redis.NewRedisService()
	rabbitServices := services_rabbitmq.NewRabbitmqService(pool)

	return &InjectorServices{
		RedisServices:  redisService,
		RabbitServices: rabbitServices,
	}
}
