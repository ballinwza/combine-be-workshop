package services_rabbitmq

import (
	"log"
	"os"

	"github.com/ballinwza/combine-be-workshop/services/mutex"
	services_redis "github.com/ballinwza/combine-be-workshop/services/redis"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitmqService struct {
	conn  *amqp091.Connection
	ch    *amqp091.Channel
	redis *services_redis.RedisService
	pool  *mutex.ClientPool
}

func NewRabbitmqService(pool *mutex.ClientPool) *RabbitmqService {
	rabbitConn := os.Getenv("RABBIT_CONN")

	conn, err := amqp091.Dial("amqp://" + rabbitConn)
	if err != nil {
		panic("Failed to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to connect channel")
	}

	redis := services_redis.NewRedisService()

	return &RabbitmqService{
		conn:  conn,
		ch:    ch,
		redis: redis,
		pool:  pool,
	}
}

func (s *RabbitmqService) Close() {
	s.ch.Close()
	s.conn.Close()
}
