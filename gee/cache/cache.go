package cache

import (
	"sync"
	"trie-demo/cache/lru"
)

type cache struct {
	lock       sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, val ByteView) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, val)
}

func (c *cache) get(key string) (val ByteView, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
