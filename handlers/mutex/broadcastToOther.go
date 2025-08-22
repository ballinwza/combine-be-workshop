package handlers_mutex

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
)

func (p *ClientPool) BroadcastToOther(message []byte, excludeUserId string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for userId, conn := range p.clients {
		if userId != excludeUserId {
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Printf("Error to broascast userId : %v\n", err)
				delete(p.clients, userId)
				conn.Close()
			}
		}
	}
}
