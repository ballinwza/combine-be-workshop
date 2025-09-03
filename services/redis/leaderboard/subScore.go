package services_redis_leaderboard

import (
	"context"
	"fmt"
)

func (s *RedisLeaderboardService) SubscribeLeaderboard(ctx context.Context) (<-chan []PlayerScore, error) {
	channel := make(chan []PlayerScore)

	sub := s.rdb.Subscribe(ctx, leaderboardChannel)
	_, err := sub.Receive(ctx)
	if err != nil {
		fmt.Printf("Error SubscribeLeaderboard : %v\n", err)
		return nil, err
	}

	go func() {
		defer close(channel)
		defer sub.Close()

		ch := sub.Channel()

		for range ch {

			newLeaderboard, err := s.GetSortScore(ctx)
			if err != nil {
				fmt.Printf("Error SubscribeLeaderboard : %v\n", err)
				continue
			}

			channel <- newLeaderboard
		}

	}()

	return channel, nil
}
