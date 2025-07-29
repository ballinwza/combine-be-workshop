package services_redis_cache

import "github.com/redis/go-redis/v9"

type RedisCache struct {
	rdb *redis.Client
}

func NewRedisCache(rdb *redis.Client) *RedisCache {
	return &RedisCache{
		rdb: rdb,
	}
}
