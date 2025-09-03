package services_redis_chat

import (
	"context"
	"fmt"
)

func (s *RedisChatService) PublishChat(ctx context.Context, message []byte) error {
	if err := s.rdb.LPush(ctx, chatKey, message).Err(); err != nil {
		fmt.Printf("Error PublishChat : %v\n", err)

		return err
	}

	err := s.rdb.Publish(ctx, chatChennel, message).Err()
	if err != nil {
		fmt.Printf("Error PublishChat : %v\n", err)
		return err

	}

	return nil
}
