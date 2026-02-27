package pokecache

import (
	"sync"
	"time"
)

// NewCache creates a new Cache and starts a background reaping loop.
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}