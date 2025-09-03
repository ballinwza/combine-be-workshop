package services_redis_leaderboard

import (
	"context"
	"fmt"
)

type PlayerScore struct {
	Username string  `json:"username"`
	Score    float64 `json:"score"`
}

func (s RedisLeaderboardService) GetSortScore(ctx context.Context) ([]PlayerScore, error) {
	scores, err := s.rdb.ZRevRangeWithScores(ctx, leaderboardDataSet, 0, 9).Result()
	if err != nil {
		fmt.Printf("Error GetSortScore : %v\n", err)
		return nil, err
	}

	leaderboard := make([]PlayerScore, len(scores))
	for index, value := range scores {
		leaderboard[index] = PlayerScore{
			Username: value.Member.(string),
			Score:    value.Score,
		}
	}

	return leaderboard, nil
}
