package handlers

import (
	services_redis "github.com/ballinwza/combine-be-workshop/services/redis"
	"github.com/gofiber/fiber/v2"
)

type LeaderboardHandler struct {
	RedisService *services_redis.RedisService
}

func (h *LeaderboardHandler) GetAllScore(c *fiber.Ctx) error {
	result, err := h.RedisService.GetAllScore(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve leaderboard",
		})
	}

	return c.JSON(result)
}
