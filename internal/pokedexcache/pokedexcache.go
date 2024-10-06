package pokedexcache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.RWMutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.RWMutex{},
	}

	go c.purgeLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	entry, ok := c.cache[key]

	return entry.val, ok
}

func (c *Cache) purgeLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.purgeCache(interval)
	}
}

func (c *Cache) purgeCache(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeAgo := time.Now().UTC().Add(-interval)
	for key, entry := range c.cache {
		if entry.createdAt.Before(timeAgo) {
			delete(c.cache, key)
		}
	}
}
