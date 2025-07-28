package services_redis_chat

import (
	"context"
	"log"
)

func (s *RedisChatService) PublishChat(ctx context.Context, message []byte) error {
	if err := s.rdb.LPush(ctx, s.chatHistoryKey, message).Err(); err != nil {
		log.Printf("บันทึก Chat ลง Redis List ไม่สำเร็จ: %v", err)
		return err
	}

	err := s.rdb.Publish(ctx, s.channel, message).Err()
	if err != nil {
		log.Printf("Publish ไปยัง Redis ไม่สำเร็จ: %v", err)
		return err

	}

	return nil
}
