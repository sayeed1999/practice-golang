package counter

import "sync"

type Counter struct {
	// mu    sync.Mutex
	sync.Mutex // embed sync.Mutex into struct
	value      int
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}
