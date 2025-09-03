package services_redis_chat

import (
	"context"
	"fmt"
)

func (s *RedisChatService) SubscribeChat(ctx context.Context) (<-chan []ChatMessage, error) {
	channel := make(chan []ChatMessage)
	sub := s.rdb.Subscribe(ctx, chatChennel)
	_, err := sub.Receive(ctx)
	if err != nil {
		fmt.Printf("Error SubscribeChat : %v\n", err)
		return nil, err
	}

	go func() {
		defer close(channel)
		defer sub.Close()

		ch := sub.Channel()

		for range ch {
			newMessage, err := s.GetChatHistory(ctx)
			if err != nil {
				fmt.Printf("Error SubscribeChat : %v\n", err)
				continue
			}
			channel <- newMessage
		}
	}()

	return channel, nil
}
