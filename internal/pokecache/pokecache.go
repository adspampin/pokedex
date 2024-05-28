package pokecache

import (
	"sync"
	"time"
)


type Cache struct{
    cacheEntries map[string]cacheEntry
    mu *sync.Mutex
}

type cacheEntry struct{
    createdAt time.Time
    val []byte
}


func NewCache(dur time.Duration) Cache{

    cache := Cache{
        cacheEntries: make(map[string]cacheEntry),
        mu: &sync.Mutex{},
    }

    go cache.reapLoop(dur)

    return cache
}

func (c Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()

    entry := cacheEntry{
        createdAt: time.Now(),
        val: val,
    }

    c.cacheEntries[key] = entry
}


func (c Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    entry, ok := c.cacheEntries[key]
    return entry.val, ok
}

func (c Cache) reapLoop(duration time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()

    t := time.Now()

    for key, entry := range c.cacheEntries {
        elapsed := t.Sub(entry.createdAt)
        if elapsed > duration {
           delete(c.cacheEntries, key)
        }
    }
}



