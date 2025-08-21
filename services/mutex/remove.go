package mutex

import "github.com/gofiber/contrib/websocket"

func (p *ClientPool) Remove(conn *websocket.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.clients, conn)
}
