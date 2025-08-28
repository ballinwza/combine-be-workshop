package services_rabbitmq_booking

import (
	handlers_mutex "github.com/ballinwza/combine-be-workshop/handlers/mutex"
	services_redis "github.com/ballinwza/combine-be-workshop/services/redis"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMqBookingService struct {
	conn  *amqp091.Connection
	ch    *amqp091.Channel
	redis *services_redis.RedisService
	pool  *handlers_mutex.ClientPool
}

func NewRabbitMqBookingService(conn *amqp091.Connection,
	ch *amqp091.Channel,
	redis *services_redis.RedisService,
	pool *handlers_mutex.ClientPool) *RabbitMqBookingService {

	return &RabbitMqBookingService{
		conn:  conn,
		ch:    ch,
		redis: redis,
		pool:  pool,
	}
}
