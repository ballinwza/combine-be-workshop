package main

import (
	"fmt"
	"log"

	"github.com/ballinwza/combine-be-workshop/configs"
	"github.com/ballinwza/combine-be-workshop/services/services_mutex"
	"github.com/joho/godotenv"

	handlers_http "github.com/ballinwza/combine-be-workshop/handlers/http"
	handlers_ws "github.com/ballinwza/combine-be-workshop/handlers/ws"
	"github.com/ballinwza/combine-be-workshop/services"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error no .env file found : %v", err)
	}

	app := fiber.New(fiber.Config{
		AppName: "Combine BE Workshop v1.0.0",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	allInjector := services.NewInjectorServices()
	pool := services_mutex.NewClientPool()

	leaderboardInjector := &handlers_http.LeaderboardHandler{
		RedisLeaderboardServices: allInjector.RedisServices.RedisLeaderboardServices,
	}

	tickerInjector := &handlers_http.TicketQueueHandler{
		RabbitmqService: allInjector.TicketQueueServices,
	}

	// Worker 1
	tickerInjector.WorkerPaymentTicket()
	defer tickerInjector.RabbitmqService.Close()

	// Worker 2
	// tickerInjector.WorkerPaymentTicket()
	// defer tickerInjector.RabbitmqService.Close()

	cacheInjector := &handlers_http.BasicCacheHandler{
		CacheService: allInjector.RedisServices.RedisCacheServices,
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world !!!")
	})

	wsGroup := app.Group("/ws", configs.SetupWebsocketConfig)
	wsGroup.Get("/chat/", websocket.New(handlers_ws.WsChat(allInjector.RedisServices.RedisChatServices)))
	wsGroup.Get("/leaderboard/", websocket.New(handlers_ws.WsLeaderboardScore(allInjector.RedisServices.RedisLeaderboardServices)))
	wsGroup.Get("/sub/queue/", websocket.New(handlers_ws.WsTicket(allInjector.RedisServices.RedisTicketServices, pool)))

	leaderboardGroup := app.Group("/leaderboard")
	leaderboardGroup.Get("/save", leaderboardInjector.SaveScoreByName)

	queueGroup := app.Group("/queue")
	queueGroup.Get("/sender", tickerInjector.BookingTicketHandler)

	cacheGroup := app.Group("/basic")
	cacheGroup.Post("/set/cache", cacheInjector.UserDetailCacheSaveHandler)
	cacheGroup.Get("/get/cache", cacheInjector.UserDetailCacheHandler)

	log.Println("Serving at localhost:8080...")
	log.Fatal(app.Listen(":8080"))
}
