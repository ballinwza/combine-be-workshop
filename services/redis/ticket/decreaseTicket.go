package services_redis_ticket

import (
	"context"
	"time"
)

func (s *RedisTicketService) DescreaseTicket(ctx context.Context) (*int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	res, err := s.CreateTicket(ctx)
	if err != nil {
		return nil, err
	}

	if res.RemainingTicket > 0 {
		res, err := s.rdb.DecrBy(ctx, s.ticketKey, 1).Result()
		if err != nil {
			return nil, err
		}
		result := int(res)
		return &result, nil
	}

	return &res.RemainingTicket, nil
}
