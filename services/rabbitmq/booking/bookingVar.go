package services_rabbitmq_booking

const (
	BOOKING_EXCHANGE       = "booking_exchange"
	BOOKING_QUEUE          = "booking_queue"
	DEAD_LETTER_EXCHANGE   = "booking_dlx"
	DEAD_LETTER_QUEUE      = "booking_dlq"
	RETRY_BOOKING_EXCHANGE = "retry_booking_exchange"
	RETRY_BOOKING_QUEUE    = "retry_booking_queue"
	MAX_RETRIES            = 3
)
