package services_redis_leaderboard

import (
	"context"
	"log"
	"time"
)

func (s *RedisLeaderboardService) SaveScoreByName(ctx context.Context, username string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := s.rdb.ZIncrBy(ctx, s.leaderboardDataSet, 1, username).Result()
	if err != nil {
		log.Printf("Error save score %v", err)
		return err
	}

	s.rdb.Publish(ctx, s.channel, "UPDATE")

	return nil
}
