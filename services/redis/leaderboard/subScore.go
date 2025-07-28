package services_redis_leaderboard

import (
	"context"
	"log"
)

func (s *RedisLeaderboardService) SubscribeLeaderboard(ctx context.Context) (<-chan []PlayerScore, error) {
	channel := make(chan []PlayerScore)

	sub := s.rdb.Subscribe(ctx, s.channel)
	_, err := sub.Receive(ctx)
	if err != nil {
		log.Printf("เกิดข้อผิดพลาดตอน Subscribe: %v", err)
		return nil, err
	}

	go func() {
		defer close(channel)
		defer sub.Close()

		ch := sub.Channel()

		for range ch {
			log.Println("ได้รับสัญญาณอัปเดต Leaderboard!")

			newLeaderboard, err := s.GetSortScore(ctx)
			if err != nil {
				log.Printf("ดึงข้อมูล Leaderboard ใหม่ไม่สำเร็จ: %v", err)
				continue
			}

			channel <- newLeaderboard
		}

	}()

	return channel, nil
}
