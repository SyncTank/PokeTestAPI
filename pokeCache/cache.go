package cache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	PokeCache map[string]cacheEntry
	lock      sync.Mutex
	internal  time.Duration
}

func NewCache(internal time.Duration) *Cache {
	newCache := &Cache{
		PokeCache: make(map[string]cacheEntry),
		lock:      sync.Mutex{},
		internal:  internal,
	}
	go newCache.reapLoop()
	return newCache
}

func (cache *Cache) AddCache(key string, val []byte) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	currentEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	if cache.PokeCache == nil {
		cache.PokeCache = make(map[string]cacheEntry)
	}
	cache.PokeCache["key"] = currentEntry
}

func (cache *Cache) GetCache(key string) ([]byte, bool) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	result, ok := cache.PokeCache[key]
	if !ok {
		fmt.Printf("Item not found: %s\n", key)
		return nil, false
	}
	return result.val, true
}

func (cache *Cache) reapLoop() {
	fmt.Println("New cache call")
	ticker := time.NewTicker(cache.internal)

	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick", t)
		}
	}
}
