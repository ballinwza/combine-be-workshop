package services_redis_chat

import (
	"context"
	"encoding/json"
	"fmt"
)

type ChatMessage struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func (s *RedisChatService) GetChatHistory(ctx context.Context) ([]ChatMessage, error) {
	history, err := s.rdb.LRange(ctx, chatKey, 0, 49).Result()
	if err != nil {
		fmt.Printf("Error GetChatHistory : %v\n", err)
		return nil, err
	}

	messages := make([]ChatMessage, 0, len(history))

	for _, msgStr := range history {
		var msg ChatMessage
		if err := json.Unmarshal([]byte(msgStr), &msg); err != nil {
			fmt.Println("Fuckig bug can't Unmarshal History idiot")
		}

		messages = append(messages, msg)
	}

	return messages, nil
}
