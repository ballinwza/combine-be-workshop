package services_redis_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v9"
)

func (s *RedisCache) GetRedisCache(ctx context.Context) {
	var wanted Test
	mycache := cache.New(&cache.Options{
		Redis:      s.rdb,
		LocalCache: cache.NewTinyLFU(10, time.Second),
	})

	if err := mycache.Get(ctx, key, &wanted); err == nil {
		fmt.Printf("✅ Value from Cache %v", wanted)
		return
	} else {
		fmt.Printf("❌ fetch new Value %v", err)
		return
	}
}
