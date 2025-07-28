package services_redis_chat

import (
	"context"
	"encoding/json"
	"log"
)

type ChatMessage struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func (s *RedisChatService) GetChatHistory(ctx context.Context) ([]ChatMessage, error) {
	history, err := s.rdb.LRange(ctx, s.chatHistoryKey, 0, 49).Result()
	if err != nil {
		return nil, err
	}

	messages := make([]ChatMessage, 0, len(history))

	for _, msgStr := range history {
		var msg ChatMessage
		if err := json.Unmarshal([]byte(msgStr), &msg); err != nil {
			log.Println("Fuckig bug can't Unmarshal History idiot")
		}

		messages = append(messages, msg)
	}

	return messages, nil
}
