package counter

import "sync"

type Counter struct {
	mu sync.Mutex
	// sync.Mutex // embed sync.Mutex into struct DONT DO THIS!
	value int
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
