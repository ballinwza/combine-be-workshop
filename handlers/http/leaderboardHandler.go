package handlers_http

import (
	services_redis "github.com/ballinwza/combine-be-workshop/services/redis"
	services_redis_leaderboard "github.com/ballinwza/combine-be-workshop/services/redis/leaderboard"
	"github.com/gofiber/fiber/v2"
)

type LeaderboardHandler struct {
	RedisService             *services_redis.RedisService
	RedisLeaderboardServices *services_redis_leaderboard.RedisLeaderboardService
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

func (h *LeaderboardHandler) SaveScoreByName(c *fiber.Ctx) error {
	username := c.Query("username")

	err := h.RedisLeaderboardServices.SaveScoreByName(c.Context(), username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not save score into leaderboard",
		})
	}

	return c.JSON(true)
}
