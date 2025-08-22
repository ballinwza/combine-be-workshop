package handlers

import (
	handlers_http "github.com/ballinwza/combine-be-workshop/handlers/http"
	handlers_mutex "github.com/ballinwza/combine-be-workshop/handlers/mutex"
	handlers_ws "github.com/ballinwza/combine-be-workshop/handlers/ws"
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
	rabbitPool := handlers_mutex.NewClientPool()

	basicCacheHandler := handlers_http.NewBasicCacheHandler()
	leaderboardHandler := handlers_http.NewLeaderboardHandler()
	wsChatHandler := handlers_ws.NewWsChatHandler()
	wsLeaderboardHandler := handlers_ws.NewWsLeaderboardHandler()

	ticketHandler := handlers_http.NewTicketHandler(rabbitPool)
	wsTicketHandler := handlers_ws.NewWsTicketHandler(rabbitPool)

	return &AllHandler{
		BasicCacheHandler:    basicCacheHandler,
		LeaderboardHandler:   leaderboardHandler,
		TicketHandler:        ticketHandler,
		WsChatHandler:        wsChatHandler,
		WsLeaderboardHandler: wsLeaderboardHandler,
		WsTicketHandler:      wsTicketHandler,
	}
}
