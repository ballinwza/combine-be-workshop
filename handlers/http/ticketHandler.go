package handlers_http

import (
	handlers_mutex "github.com/ballinwza/combine-be-workshop/handlers/mutex"
	"github.com/ballinwza/combine-be-workshop/services"
	services_rabbitmq "github.com/ballinwza/combine-be-workshop/services/rabbitmq"
	"github.com/gofiber/fiber/v2"
)

type TicketHandler struct {
	RabbitmqService *services_rabbitmq.RabbitmqService
}

func NewTicketHandler(pool *handlers_mutex.ClientPool) TicketHandler {
	rabbitmqService := services.NewInjectorServices(pool).RabbitServices
	return TicketHandler{
		RabbitmqService: rabbitmqService,
	}
}

func (h *TicketHandler) BookingTicketHandler(c *fiber.Ctx) error {
	userId := c.Params("userId")
	err := h.RabbitmqService.BookingTicket(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed send payment ticket",
		})
	}

	return c.Status(fiber.StatusOK).JSON(c.JSON(fiber.Map{
		"success": "sended payment queue",
	}))
}

func (h *TicketHandler) WorkerPaymentTicket() {
	go h.RabbitmqService.WorkerPaymentTicket()
}
