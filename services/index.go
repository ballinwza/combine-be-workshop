package services

import "github.com/gofiber/fiber/v2"

type Services struct {
	ctx *fiber.Ctx
}

func NewServices(ctx *fiber.Ctx) *Services {
	return &Services{
		ctx: ctx,
	}
}
