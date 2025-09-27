package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]CacheEntry
	mu   sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheEntry{data: val, createdAt: time.Now()}
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return v.data, true
}

func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		for k, v := range c.data {
			if time.Since(v.createdAt) >= interval {
				c.mu.Lock()
				delete(c.data, k)
				c.mu.Unlock()
			}
		}
	}

}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		data: make(map[string]CacheEntry),
	}
	go cache.readLoop(interval)
	return &cache
}
