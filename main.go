package main

import (
	"fmt"
	"log"

	"github.com/ballinwza/combine-be-workshop/configs"
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

	// Worker
	allInjector.TicketQueueServices.WorkerPaymentTicket()

	leaderboardInjector := &handlers_http.LeaderboardHandler{
		RedisLeaderboardServices: allInjector.RedisServices.RedisLeaderboardServices,
	}

	tickerInjector := &handlers_http.TicketQueueHandler{
		RabbitmqService: allInjector.TicketQueueServices,
	}
	cacheInjector := &handlers_http.BasicCacheHandler{
		CacheService: allInjector.RedisServices.RedisCacheServices,
	}

	app.Use("/ws", configs.SetupWebsocketConfig)

	/*
		app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
			log.Println(c.Locals("allowed"))  // true
			log.Println(c.Params("id"))       // 123
			log.Println(c.Query("v"))         // 1.0
			log.Println(c.Cookies("session")) //

			var (
				mt  int
				msg []byte
				err error
			)

			for {
				if mt, msg, err = c.ReadMessage(); err != nil {
					log.Println("read:", err)
					break
				}
				log.Printf("recv: %s", msg)

				if err = c.WriteMessage(mt, msg); err != nil {
					log.Println("write:", err)
					break
				}
			}
		}))
	*/

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Get("/ws/chat/", websocket.New(handlers_ws.WsChat(allInjector.RedisServices.RedisChatServices)))
	app.Get("/ws/leaderboard/", websocket.New(handlers_ws.WsLeaderboardScore(allInjector.RedisServices.RedisLeaderboardServices)))

	app.Get("/leaderboard/save", leaderboardInjector.SaveScoreByName)

	app.Get("/queue/sender", tickerInjector.BookingTicket)

	app.Post("/basic/set/cache", cacheInjector.UserDetailCacheSaveHandler)
	app.Get("/basic/get/cache", cacheInjector.UserDetailCacheHandler)

	log.Println("Serving at localhost:8080...")
	log.Fatal(app.Listen(":8080"))
}
