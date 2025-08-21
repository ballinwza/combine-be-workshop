package mutex

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
)

func (p *ClientPool) Broadcast(message []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for conn := range p.clients {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Printf("Error writing message : %v\n", err)
			conn.Close()
			delete(p.clients, conn)
		}
	}
}
