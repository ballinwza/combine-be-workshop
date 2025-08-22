package handlers_mutex

import "github.com/gofiber/contrib/websocket"

func (p *ClientPool) Add(userId string, conn *websocket.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.clients[userId] = conn
}
