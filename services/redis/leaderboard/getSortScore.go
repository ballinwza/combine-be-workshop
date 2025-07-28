package services_redis_leaderboard

import "context"

type PlayerScore struct {
	Username string  `json:"username"`
	Score    float64 `json:"score"`
}

func (s RedisLeaderboardService) GetSortScore(ctx context.Context) ([]PlayerScore, error) {
	scores, err := s.rdb.ZRevRangeWithScores(ctx, s.leaderboardDataSet, 0, 9).Result()
	if err != nil {
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
