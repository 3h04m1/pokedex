package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	// add a mutex here
	mut  sync.Mutex
	data map[string]CacheEntry
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) []byte {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.data[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	return val
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	val, ok := c.data[key]
	return val.val, ok
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data: make(map[string]CacheEntry),
	}
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.readLoop(
				interval,
			)
		}
	}()
	return c
}

func (c *Cache) readLoop(interval time.Duration) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for key, entry := range c.data {
		if time.Since(entry.createdAt) > interval {
			delete(c.data, key)
		}
	}
}
