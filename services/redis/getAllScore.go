package services_redis

import (
	"context"
	"time"
)

func (c *RedisService) GetAllScore(ctx context.Context) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res1, err := c.rdb.Get(ctx, "bike:1").Result()
	if res1 == "" {
		value := "Deimos"
		c.rdb.Set(ctx, "bike:1", value, 0)
		return &value, nil
	}

	if err != nil {
		return nil, err
	}

	return &res1, nil

}
