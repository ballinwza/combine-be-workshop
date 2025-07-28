package services_redis_leaderboard

import "github.com/redis/go-redis/v9"

type RedisLeaderboardService struct {
	rdb                *redis.Client
	channel            string
	leaderboardDataSet string
}

func NewRedisLeaderboardService(rdb *redis.Client) *RedisLeaderboardService {
	const leaderboardChannel = "leaderboard_channel"
	const leaderboardDataSet = "leaderboard_score"

	return &RedisLeaderboardService{
		rdb:                rdb,
		channel:            leaderboardChannel,
		leaderboardDataSet: leaderboardDataSet,
	}
}
