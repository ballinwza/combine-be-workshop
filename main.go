package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ballinwza/combine-be-workshop/configs"
	"github.com/ballinwza/combine-be-workshop/handlers"
	"github.com/joho/godotenv"

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

	allHandler := handlers.NewAllHandler()

	// Setup Exchange, Queue
	err = allHandler.TicketHandler.RabbitmqService.RabbitbookingService.SetupBooking()
	if err != nil {
		log.Fatalf("Failed to setup RabbitMQ queues: %v", err)
	}

	err = allHandler.TicketHandler.RabbitmqService.RabbitbookingService.SetupBookingDeadLetter()
	if err != nil {
		log.Fatalf("Failed to setup RabbitMQ queues: %v", err)
	}

	err = allHandler.TicketHandler.RabbitmqService.RabbitbookingService.SetupBookingRetry()
	if err != nil {
		log.Fatalf("Failed to setup RabbitMQ retry queue: %v", err)
	}

	// Worker 1
	allHandler.TicketHandler.WorkerPaymentTicket()
	// defer allHandler.TicketHandler.RabbitmqService.Close()

	// Worker 2
	// tickerInjector.WorkerPaymentTicket()
	// defer tickerInjector.RabbitmqService.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world !!!")
	})

	wsGroup := app.Group("/ws", configs.SetupWebsocketConfig)
	wsGroup.Get("/chat/", websocket.New(allHandler.WsChatHandler.WsChat()))
	wsGroup.Get("/leaderboard/", websocket.New(allHandler.WsLeaderboardHandler.WsLeaderboardScore()))
	wsGroup.Get("/sub/queue/:userId", websocket.New(allHandler.WsTicketHandler.WsTicket()))

	leaderboardGroup := app.Group("/leaderboard")
	leaderboardGroup.Get("/save", allHandler.LeaderboardHandler.SaveScoreByName)

	queueGroup := app.Group("/queue")
	queueGroup.Get("/sender/:userId", allHandler.TicketHandler.BookingTicketHandler)

	cacheGroup := app.Group("/basic")
	cacheGroup.Post("/set/cache", allHandler.BasicCacheHandler.UserDetailCacheSaveHandler)
	cacheGroup.Get("/get/cache", allHandler.BasicCacheHandler.UserDetailCacheHandler)

	// log.Println("Serving at localhost:8080...")
	// log.Fatal(app.Listen(":8080"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		log.Println("Serving at localhost:8080...")
		if err := app.Listen(":8080"); err != nil {
			log.Panic(err)
		}
	}()

	<-quit

	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	if err := app.Shutdown(); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	}

	allHandler.TicketHandler.RabbitmqService.Close() // ปิด Connection ตอนจบโปรแกรม
	log.Println("Server shut down.")
}
