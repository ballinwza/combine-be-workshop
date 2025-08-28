package services_rabbitmq_booking

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func (s *RabbitMqBookingService) SetupBooking() error {
	// Booking
	err := s.ch.ExchangeDeclare(BOOKING_EXCHANGE, "fanout", true, false, false, false, nil)
	if err != nil {
		return err
	}

	args := amqp091.Table{
		"x-dead-letter-exchange":    DEAD_LETTER_EXCHANGE,
		"x-dead-letter-routing-key": DEAD_LETTER_QUEUE, // ระบุ key ที่จะใช้ใน DLX
	}

	_, err = s.ch.QueueDeclare(
		BOOKING_QUEUE,
		true,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		return err
	}

	err = s.ch.QueueBind(BOOKING_QUEUE, "", BOOKING_EXCHANGE, false, nil)
	if err != nil {
		return fmt.Errorf("failed to bind main queue to main exchange: %w", err)
	}

	fmt.Println("RabbitMQ queues and DLQ setup successfully.")
	return nil
}
