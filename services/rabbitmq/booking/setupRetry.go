package services_rabbitmq_booking

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func (s *RabbitMqBookingService) SetupBookingRetry() error {
	err := s.ch.ExchangeDeclare(RETRY_BOOKING_EXCHANGE, "direct", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("failed to declare retry exchange: %w", err)
	}

	args := amqp091.Table{
		"x-message-ttl":             2000, // 10,000 millisec
		"x-dead-letter-exchange":    BOOKING_EXCHANGE,
		"x-dead-letter-routing-key": BOOKING_QUEUE,
	}
	_, err = s.ch.QueueDeclare(RETRY_BOOKING_QUEUE, true, false, false, false, args)
	if err != nil {
		return fmt.Errorf("failed to declare retry queue: %w", err)
	}

	err = s.ch.QueueBind(RETRY_BOOKING_QUEUE, RETRY_BOOKING_QUEUE, RETRY_BOOKING_EXCHANGE, false, nil)
	if err != nil {
		return fmt.Errorf("failed to declare retry queue: %w", err)
	}

	return nil
}
