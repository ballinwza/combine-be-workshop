package services_rabbitmq_booking

import (
	"fmt"
)

func (s *RabbitMqBookingService) SetupBookingDeadLetter() error {
	// Dead Letter
	err := s.ch.ExchangeDeclare(
		DEAD_LETTER_EXCHANGE,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	_, err = s.ch.QueueDeclare(
		DEAD_LETTER_QUEUE,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = s.ch.QueueBind(DEAD_LETTER_QUEUE, DEAD_LETTER_QUEUE, DEAD_LETTER_EXCHANGE, false, nil)
	if err != nil {
		return err
	}

	fmt.Println("RabbitMQ queues and DLQ setup successfully.")
	return nil
}
