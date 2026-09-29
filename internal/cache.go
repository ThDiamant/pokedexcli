package internal

import (
	"sync"
	"time"
)

type Cache struct {
	contents map[string]cacheEntry
	mu       sync.Mutex
	ticker   time.Ticker
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(time.Second)
	newCache := Cache{
		contents: map[string]cacheEntry{},
		mu:       sync.Mutex{},
		ticker:   *ticker,
	}
	defer ticker.Stop()

	go newCache.readLoop(interval)
	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.contents[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.contents[key]
	c.mu.Unlock()
	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) readLoop(interval time.Duration) {
	for range c.ticker.C {
		for k, entry := range c.contents {
			if time.Since(entry.createdAt) > interval {
				delete(c.contents, k)
			}
		}
	}
}
