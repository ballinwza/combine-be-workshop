package handlers_mutex

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
)

func (p *ClientPool) SendToUser(userId string, message []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if conn, ok := p.clients[userId]; ok {
		fmt.Printf("Sending to user UserID: %s\n", userId)
		err := conn.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("Client disconnected : %v\n", err)
			} else {

				fmt.Printf("Error writing message : %v\n", err)
			}
			delete(p.clients, userId)
			conn.Close()
		}
	} else {
		fmt.Printf("No active WebSocket connection found for UserID: %s\n", userId)
	}
}
