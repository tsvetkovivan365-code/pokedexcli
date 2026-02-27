package pokecache

import (
	"sync"
	"time"
)

// cacheEntry represents a single entry in the cache.
type cacheEntry struct {
	createdAt time.Time
	val []byte
}

// Cache represents a cache that stores entries by string keys.
type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.Mutex
}