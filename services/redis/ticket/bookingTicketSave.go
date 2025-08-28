package services_redis_ticket

import (
	"context"
	"encoding/json"
)

func (s *RedisTicketService) SaveBookingTicket(ctx context.Context, ticketId string, userId string, isSuccess bool, isPending bool, message string, statusType BookingStatus) (*BookingTicket, error) {
	redisKey := "booking_ticket:" + userId

	payload := BookingTicket{
		Id:        &ticketId,
		UserId:    userId,
		IsSuccess: isSuccess,
		IsPending: isPending,
		Message:   message,
		Type:      statusType,
	}

	jsonData, _ := json.Marshal(payload)
	err := s.rdb.RPush(ctx, redisKey, jsonData).Err()
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
