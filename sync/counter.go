package sync

import "sync"

type Counter struct {
	state int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.state++
}

func (c *Counter) Value() int {
	return c.state
}

func NewCounter() *Counter {
	return &Counter{}
}
