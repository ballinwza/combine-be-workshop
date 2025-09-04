package handlers

import (
	"os"

	handlers_http "github.com/ballinwza/combine-be-workshop/handlers/http"
	handlers_mutex "github.com/ballinwza/combine-be-workshop/handlers/mutex"
	handlers_ws "github.com/ballinwza/combine-be-workshop/handlers/ws"
	"github.com/ballinwza/combine-be-workshop/services"
	"github.com/redis/go-redis/v9"
)

type AllHandler struct {
	BasicCacheHandler    handlers_http.BasicCacheHandler
	LeaderboardHandler   handlers_http.LeaderboardHandler
	TicketHandler        handlers_http.TicketHandler
	WsChatHandler        handlers_ws.WsChatHandler
	WsLeaderboardHandler handlers_ws.WsLeaderboardHandler
	WsTicketHandler      handlers_ws.WsTicketHandler
}

func NewAllHandler() *AllHandler {
	// Setup Pool
	rabbitPool := handlers_mutex.NewClientPool()

	// Setup Redis
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	// Setup service
	allService := services.NewInjectorServices(rabbitPool, rdb)

	// Setup handler
	basicCacheHandler := handlers_http.NewBasicCacheHandler(*allService)
	leaderboardHandler := handlers_http.NewLeaderboardHandler(*allService)
	wsChatHandler := handlers_ws.NewWsChatHandler(*allService)
	wsLeaderboardHandler := handlers_ws.NewWsLeaderboardHandler(*allService)

	ticketHandler := handlers_http.NewTicketHandler(rabbitPool, *allService)
	wsTicketHandler := handlers_ws.NewWsTicketHandler(rabbitPool, *allService)

	return &AllHandler{
		BasicCacheHandler:    basicCacheHandler,
		LeaderboardHandler:   leaderboardHandler,
		TicketHandler:        ticketHandler,
		WsChatHandler:        wsChatHandler,
		WsLeaderboardHandler: wsLeaderboardHandler,
		WsTicketHandler:      wsTicketHandler,
	}
}
