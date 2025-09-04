package services_rabbitmq

import (
	"fmt"
	"log"
	"os"

	handlers_mutex "github.com/ballinwza/combine-be-workshop/handlers/mutex"
	services_rabbitmq_booking "github.com/ballinwza/combine-be-workshop/services/rabbitmq/booking"
	services_redis "github.com/ballinwza/combine-be-workshop/services/redis"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitmqService struct {
	conn                 *amqp091.Connection
	ch                   *amqp091.Channel
	redis                *services_redis.RedisService
	pool                 *handlers_mutex.ClientPool
	RabbitbookingService services_rabbitmq_booking.RabbitMqBookingService
}

func NewRabbitmqService(pool *handlers_mutex.ClientPool, rdb *services_redis.RedisService) *RabbitmqService {
	rabbitConn := os.Getenv("RABBIT_CONN")

	conn, err := amqp091.Dial("amqp://" + rabbitConn)
	if err != nil {
		fmt.Printf("Failed to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to connect channel")
	}

	rabbitbookingService := services_rabbitmq_booking.NewRabbitMqBookingService(conn, ch, rdb, pool)

	return &RabbitmqService{
		conn:                 conn,
		ch:                   ch,
		redis:                rdb,
		pool:                 pool,
		RabbitbookingService: *rabbitbookingService,
	}
}

func (s *RabbitmqService) Close() {
	s.ch.Close()
	s.conn.Close()
}
