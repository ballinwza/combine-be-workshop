package services_redis_ticket

type BookingTicket struct {
	Id        *string       `json:"Id,omitempty" redis:"id,omitempty"`
	UserId    string        `json:"UserId" redis:"userId"`
	IsSuccess bool          `json:"IsSuccess" redis:"isSuccess"`
	IsPending bool          `json:"IsPending" redis:"isPending"`
	Message   string        `json:"Message" redis:"message"`
	Type      BookingStatus `json:"Type" redis:"type"`
}

type NotificationTicket struct {
	IsSuccess bool          `json:"IsSuccess" redis:"IsSuccess"`
	Message   string        `json:"Message" redis:"message"`
	Type      BookingStatus `json:"Type" redis:"type"`
}

type RemainingTicket struct {
	RemainingTicket int           `json:"RemainingTicket" redis:"remainingTicket"`
	Type            BookingStatus `json:"Type" redis:"type"`
}

type BookingStatus string

const (
	Pending      BookingStatus = "booking_pending"
	Successed    BookingStatus = "booking_successed"
	Failed       BookingStatus = "booking_failed"
	Remaining    BookingStatus = "ticket_remaining"
	Notification BookingStatus = "ticket_notification"
)

const (
	ticketKey     = "available_ticket"
	ticketChannel = "ticket_remaining_channel"
)
