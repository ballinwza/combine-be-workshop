package main

import (
	"github.com/ballinwza/combine-be-workshop/handlers"
	"github.com/ballinwza/combine-be-workshop/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Combine BE Workshop v1.0.0",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Context-Type, Accept",
	}))

	injector := services.NewInjectorServices()

	leaderboardHandler := &handlers.LeaderboardHandler{
		RedisService: injector.RedisService,
	}
	app.Get("/leaderboard", leaderboardHandler.GetAllScore)

	app.Listen(":3001")
}
