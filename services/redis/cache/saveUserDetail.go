package services_redis_cache

import (
	"context"
	"fmt"
	"time"
)

type UserDetail struct {
	Name       string `redis:"name" json:"name"`
	Age        uint32 `redis:"age" json:"age"`
	ExpireTime uint32 `json:"expireTime"`
}

func (s *RedisCache) SaveUserDetailCache(ctx context.Context, key string, userDetail UserDetail, cacheTime uint32) (bool, error) {
	pipe := s.rdb.Pipeline()

	pipe.HSet(ctx, key, userDetail)

	if cacheTime != 0 {
		fmt.Printf("%v", cacheTime)
		duration := time.Duration(userDetail.ExpireTime) * time.Second
		pipe.Expire(ctx, key, duration)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return false, err
	}

	return true, nil
}
