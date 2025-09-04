package handlers_http

import (
	"github.com/ballinwza/combine-be-workshop/services"
	services_redis_leaderboard "github.com/ballinwza/combine-be-workshop/services/redis/leaderboard"
	"github.com/gofiber/fiber/v2"
)

type LeaderboardHandler struct {
	RedisLeaderboardServices *services_redis_leaderboard.RedisLeaderboardService
}

func NewLeaderboardHandler(allService services.InjectorServices) LeaderboardHandler {
	redisLeaderboardServices := allService.RedisServices.RedisLeaderboardServices

	return LeaderboardHandler{
		RedisLeaderboardServices: redisLeaderboardServices,
	}
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
