package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache() *Cache {
	return &Cache{data: map[string]cacheEntry{}, mu: &sync.Mutex{}}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.data[key] = cEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cEntry, ok := c.data[key]; ok {
		return cEntry.val, true
	}
	return nil, false
}
