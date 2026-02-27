package pokecache

import "time"

// Add adds a new entry to the cache.
func (c *Cache) Add(url string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[url] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}