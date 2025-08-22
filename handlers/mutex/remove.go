package handlers_mutex

func (p *ClientPool) Remove(userId string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.clients[userId]; ok {
		delete(p.clients, userId)
	}
}
