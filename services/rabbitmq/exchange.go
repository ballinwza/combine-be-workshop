package services_rabbitmq

import "fmt"

func (s *RabbitmqService) exchangeBookingDeclare() error {
	err := s.ch.ExchangeDeclare(
		"ticket",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Error exchangeBookingDeclare : %s\n", err)
		return err
	}

	return nil
}
