package services_redis_ticket

import (
	"context"
	"fmt"
	"time"
)

func (s *RedisTicketService) DescreaseTicket(ctx context.Context) (*int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	res, err := s.SaveRemainingTicket(ctx)
	if err != nil {
		fmt.Printf("Error DescreaseTicket : %v\n", err)
		return nil, err
	}

	if res.RemainingTicket > 0 {
		res, err := s.rdb.DecrBy(ctx, ticketKey, 1).Result()
		if err != nil {
			fmt.Printf("Error DescreaseTicket : %v\n", err)
			return nil, err
		}
		result := int(res)
		return &result, nil
	}

	return &res.RemainingTicket, nil
}
