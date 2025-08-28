package services_redis_ticket

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func (s *RedisTicketService) SaveRemainingTicket(ctx context.Context) (*RemainingTicket, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var result RemainingTicket
	chkValue, err := s.rdb.Get(ctx, s.ticketKey).Result()
	fmt.Printf("num : %v\n", chkValue)
	if chkValue == "" || err == redis.Nil {
		_, err := s.rdb.Set(ctx, s.ticketKey, 10, -1).Result()
		if err != nil {
			return nil, err
		}

		result = RemainingTicket{
			RemainingTicket: 10,
		}

		return &result, nil
	} else if err != nil {
		fmt.Printf("Error can't get remaining ticket : %v\n", err)
		return nil, err
	} else {
		convToNum, err := strconv.Atoi(chkValue)
		if err != nil {
			fmt.Printf("Convert string to int failed at GetTicket\n")
		}

		result = RemainingTicket{
			RemainingTicket: convToNum,
		}

		return &result, nil
	}

}
