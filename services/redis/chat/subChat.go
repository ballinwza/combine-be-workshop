package services_redis_chat

import (
	"context"
	"log"
)

func (s *RedisChatService) SubscribeChat(ctx context.Context) (<-chan []ChatMessage, error) {
	channel := make(chan []ChatMessage)
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
			newMessage, err := s.GetChatHistory(ctx)
			if err != nil {
				log.Printf("ดึงข้อมูล Chat ใหม่ไม่สำเร็จ: %v", err)
				continue
			}
			channel <- newMessage
		}
	}()

	return channel, nil
}
