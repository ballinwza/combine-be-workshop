package services_redis_leaderboard

import (
	"context"
	"fmt"
)

func (s *RedisLeaderboardService) SaveScoreByName(ctx context.Context, username string) error {
	_, err := s.rdb.ZIncrBy(ctx, leaderboardDataSet, 1, username).Result()
	if err != nil {
		fmt.Printf("Error SaveScoreByName : %v\n", err)
		return err
	}

	s.rdb.Publish(ctx, leaderboardChannel, "UPDATE")

	return nil
}
