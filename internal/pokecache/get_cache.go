package pokecache

// Get retrieves an entry from the cache for the given key.
func (c *Cache) Get(key string) ([]byte, bool){

	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.cache[key]

	if !ok {
		return nil, false
	}

	return entry.val, true
}