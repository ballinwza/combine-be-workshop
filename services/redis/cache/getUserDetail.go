package services_redis_cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (s *RedisCache) GetUserDetailCache(ctx context.Context, key string) (*UserDetail, error) {
	findKey, err := s.rdb.Exists(ctx, key).Result()
	if err != nil {
		fmt.Printf("❌ Error checking existence: %v\n", err)
		return nil, err
	}

	if findKey == 0 {
		fmt.Printf("🟡 Key '%s' does not exist.\n", key)
		return nil, redis.Nil
	}

	var user UserDetail
	err = s.rdb.HGetAll(ctx, key).Scan(&user)
	if err != nil {

		fmt.Printf("❌ must fetch new Value %v\n", err)
		return nil, err
	}

	fmt.Printf("✅ Value from Cache %v\n", user)
	return &user, nil
}
