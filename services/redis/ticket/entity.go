package services_redis_ticket

type NotificationTicket struct {
	UserId    string `json:"UserId" redis:"userId"`
	Status    bool   `json:"Status" redis:"status"`
	IsPending bool   `json:"IsPending" redis:"isPending"`
	Message   string `json:"Message" redis:"message"`
	Type      string `json:"Type" redis:"type"`
}

type RemainingTicket struct {
	RemainingTicket int    `json:"RemainingTicket" redis:"remainingTicket"`
	Type            string `json:"Type" redis:"type"`
}
