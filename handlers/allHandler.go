package handlers

import (
	handlers_http "github.com/ballinwza/combine-be-workshop/handlers/http"
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
	basicCacheHandler := handlers_http.NewBasicCacheHandler()
	leaderboardHandler := handlers_http.NewLeaderboardHandler()
	ticketHandler := handlers_http.NewTicketHandler()
	wsChatHandler := handlers_ws.NewWsChatHandler()
	wsLeaderboardHandler := handlers_ws.NewWsLeaderboardHandler()
	wsTicketHandler := handlers_ws.NewWsTicketHandler()

	return &AllHandler{
		BasicCacheHandler:    basicCacheHandler,
		LeaderboardHandler:   leaderboardHandler,
		TicketHandler:        ticketHandler,
		WsChatHandler:        wsChatHandler,
		WsLeaderboardHandler: wsLeaderboardHandler,
		WsTicketHandler:      wsTicketHandler,
	}
}
