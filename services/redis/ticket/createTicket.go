package services_redis_ticket

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func (s *RedisTicketService) CreateTicket(ctx context.Context) (*RemainingTicket, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var result RemainingTicket
	chkValue, err := s.rdb.Get(ctx, s.ticketKey).Result()
	if chkValue == "" {
		ticket, err := s.rdb.Set(ctx, s.ticketKey, 2, time.Minute*60).Result()
		if err != nil {
			return nil, err
		}

		convToNum, err := strconv.Atoi(ticket)
		if err != nil {
			fmt.Printf("Convert string to int failed at GetTicket")
		}

		result = RemainingTicket{
			RemainingTicket: convToNum,
		}

		return &result, nil
	}

	if err != nil {
		fmt.Printf("Error can't get remaining ticket : %v\n", err)
		return nil, err
	}

	convToNum, err := strconv.Atoi(chkValue)
	if err != nil {
		fmt.Printf("Convert string to int failed at GetTicket")
	}

	result = RemainingTicket{
		RemainingTicket: convToNum,
	}

	return &result, nil
}
