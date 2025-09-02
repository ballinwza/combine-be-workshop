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
	chkValue, err := s.rdb.Get(ctx, ticketKey).Result()

	if chkValue == "" || err == redis.Nil {
		_, err := s.rdb.Set(ctx, ticketKey, 10, -1).Result()
		if err != nil {
			fmt.Printf("Error SaveRemainingTicket : %v\n", err)
			return nil, err
		}

		result = RemainingTicket{
			RemainingTicket: 10,
		}

		return &result, nil
	} else if err != nil {
		fmt.Printf("Error SaveRemainingTicket : %v\n", err)
		return nil, err
	} else {
		convToNum, _ := strconv.Atoi(chkValue)
		result = RemainingTicket{
			RemainingTicket: convToNum,
		}

		return &result, nil
	}

}
