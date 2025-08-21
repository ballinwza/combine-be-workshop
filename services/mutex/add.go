package mutex

import "github.com/gofiber/contrib/websocket"

func (p *ClientPool) Add(conn *websocket.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.clients[conn] = true
}
