package handlers_mutex

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type ClientPool struct {
	clients map[string]*websocket.Conn
	mu      sync.Mutex
}

func NewClientPool() *ClientPool {
	return &ClientPool{
		clients: make(map[string]*websocket.Conn),
	}
}
