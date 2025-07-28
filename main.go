package main

import (
	"log"

	"github.com/ballinwza/combine-be-workshop/configs"

	handlers_http "github.com/ballinwza/combine-be-workshop/handlers/http"
	handlers_ws "github.com/ballinwza/combine-be-workshop/handlers/ws"
	"github.com/ballinwza/combine-be-workshop/services"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Combine BE Workshop v1.0.0",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	allInjector := services.NewInjectorServices()

	leaderboardInjector := &handlers_http.LeaderboardHandler{
		RedisService:             allInjector.RedisService,
		RedisLeaderboardServices: allInjector.RedisLeaderboardServices,
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

	app.Get("/ws/chat/", websocket.New(handlers_ws.WsChat(allInjector.RedisChatServices)))
	app.Get("/ws/leaderboard/", websocket.New(handlers_ws.WsLeaderboardScore(allInjector.RedisLeaderboardServices)))

	app.Get("/leaderboard/save", leaderboardInjector.SaveScoreByName)

	log.Println("Serving at localhost:3001...")
	log.Fatal(app.Listen(":3001"))
}
