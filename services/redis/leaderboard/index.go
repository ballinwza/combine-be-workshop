package services_redis_leaderboard

import "github.com/redis/go-redis/v9"

type RedisLeaderboardService struct {
	rdb *redis.Client
}

func NewRedisLeaderboardService(rdb *redis.Client) *RedisLeaderboardService {

	return &RedisLeaderboardService{
		rdb: rdb,
	}
}
