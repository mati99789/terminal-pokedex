package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu *sync.RWMutex
	v  map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		v:  make(map[string]cacheEntry),
		mu: &sync.RWMutex{},
	}

	go c.reapLoop(interval)

	return &c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.v {
			if time.Since(entry.createdAt) > interval {
				delete(c.v, key)
			}
		}

		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cache := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.v[key] = cache

}

func (c *Cache) Get(key string) (val []byte, isPresent bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.v[key]

	if !ok {
		return nil, false
	}

	return entry.val, true

}
