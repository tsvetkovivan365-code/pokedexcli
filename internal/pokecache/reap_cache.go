package pokecache

import "time"

// reapLoop deletes entries older than the given interval.
func (c *Cache) reapLoop(interval time.Duration) {
	
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mutex.Lock()
		for k, v := range c.cache {
			elapsed := time.Since(v.createdAt)
			if elapsed > interval {
				delete(c.cache, k)
			}
		}
		c.mutex.Unlock()
	}
}