package aibrain

import (
	"sync"
	"time"
)

type cache struct {
	m   sync.Map
	ttl time.Duration
}

type entry struct {
	threadID   string
	expireTime time.Time
}

func newCache(ttl time.Duration) *cache {
	return &cache{
		m:   sync.Map{},
		ttl: ttl,
	}
}

func (c *cache) Set(messageID, threadID string) {
	c.m.Store(messageID, &entry{
		threadID:   threadID,
		expireTime: time.Now().Add(c.ttl),
	})
}

func (c *cache) Get(messageID string) (string, bool) {
	item, ok := c.m.Load(messageID)
	if !ok {
		return "", false
	}

	e, ok := item.(*entry)
	if !ok || time.Now().After(e.expireTime) {
		c.m.Delete(messageID)
		return "", false
	}

	return e.threadID, true
}

func (c *cache) Delete(messageID string) {
	c.m.Delete(messageID)
}

func (c *cache) Cleanup() {
	ticker := time.NewTicker(c.ttl / 16)
	defer ticker.Stop()

	for range ticker.C {
		c.m.Range(func(k interface{}, v interface{}) bool {
			e, ok := v.(*entry)
			if !ok || time.Now().After(e.expireTime) {
				c.m.Delete(k)
			}
			return true
		})
	}
}
