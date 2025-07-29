package services_redis_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v9"
)

type Test struct {
	Name string
	Age  uint32
}

const key = "redis-cache-key"

func (s *RedisCache) SaveIntoCache(ctx context.Context) {
	mycache := cache.New(&cache.Options{
		Redis:      s.rdb,
		LocalCache: cache.NewTinyLFU(10, time.Second),
	})

	obj := &Test{
		Name: "Test Cache",
		Age:  21,
	}

	if err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: obj,
		TTL:   time.Second,
	}); err != nil {
		panic(err)
	}

	var wanted Test
	if err := mycache.Get(ctx, key, &wanted); err == nil {
		fmt.Printf("✅ Value from Cache %v", wanted)
		return
	} else {
		fmt.Printf("❌ fetch new Value %v", err)
		return
	}
}
