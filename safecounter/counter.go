package safecounter

import "sync"

type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string){
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Val(key string) int{
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}