package pokecache

import (
	"sync"
	"time"
)

type Cache struct {		//Cache struct for holding data
	Cached	map[string]cacheEntry
	mu		sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(interval time.Duration) *Cache {	//Creates a new cache of data
	c := Cache{
		Cached: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)		//Starts reapLoop goroutine to continuously clear cache
	return &c
}

func (c *Cache) reapLoop(interval time.Duration) {	//Clears cache of any data stored longer than the given interval
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		c.mu.Lock()
		for key, cache := range c.Cached {
			if (time.Since(cache.createdAt)) > interval {
				delete(c.Cached, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, value []byte) {	//Adds cacheEntry to data cache
	c.mu.Lock()
	c.Cached[key] = cacheEntry{createdAt: time.Now(), val: value}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {	//Gets data from the cache
	c.mu.Lock()
	defer c.mu.Unlock()
	if value, ok := c.Cached[key]; !ok {
		return nil, false
	} else {
		return value.val, true
	}
}