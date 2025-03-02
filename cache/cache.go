package cache

import "sync"

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewCache() *Cache{
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) (interface{}, bool){
		c.mu.RLock()
		defer c.mu.RUnlock()
		val, exists := c.data[key]
		return val, exists
}

func (c *Cache) Set(key string, value interface{}){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}