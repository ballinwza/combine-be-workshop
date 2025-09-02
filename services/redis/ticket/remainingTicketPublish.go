package services_redis_ticket

import (
	"context"
	"fmt"
)

func (s *RedisTicketService) PublishRemainingTicket(ctx context.Context, msg []byte) error {
	err := s.rdb.Publish(ctx, ticketChannel, msg).Err()
	if err != nil {
		fmt.Printf("Error PublishRemainingTicket PUBLISH : %v\n", err)
		return err
	}

	return nil
}
