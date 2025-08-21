package mutex

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type ClientPool struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

func NewClientPool() *ClientPool {
	return &ClientPool{
		clients: make(map[*websocket.Conn]bool),
	}
}
