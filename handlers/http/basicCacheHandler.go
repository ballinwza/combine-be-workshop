package handlers_http

import (
	"github.com/ballinwza/combine-be-workshop/services"
	services_redis_cache "github.com/ballinwza/combine-be-workshop/services/redis/cache"
	"github.com/gofiber/fiber/v2"
)

type BasicCacheHandler struct {
	CacheService *services_redis_cache.RedisCache
}

func NewBasicCacheHandler() BasicCacheHandler {
	redisCacheServices := services.NewInjectorServices(nil).RedisServices.RedisCacheServices
	return BasicCacheHandler{
		CacheService: redisCacheServices,
	}
}

func (h *BasicCacheHandler) UserDetailCacheHandler(c *fiber.Ctx) error {
	n := c.Query("name")

	value, err := h.CacheService.GetUserDetailCache(c.Context(), n)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed get cache",
		})
	}

	return c.Status(fiber.StatusOK).JSON(value)
}

func (h *BasicCacheHandler) UserDetailCacheSaveHandler(c *fiber.Ctx) error {
	var user services_redis_cache.UserDetail

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed get cache " + err.Error(),
		})
	}

	value, err := h.CacheService.SaveUserDetailCache(c.Context(), user.Name, user, user.ExpireTime)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed save cache",
		})
	}

	if !value {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error ",
		})
	}

	return c.Status(fiber.StatusOK).JSON(value)
}
