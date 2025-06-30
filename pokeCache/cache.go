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
}

func NewCache(internal time.Duration) Cache {

	reapLoop()
	return Cache{}
}

func AddCache(key string, val []byte) {
	cache := Cache{}
	currentEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.PokeCache["key"] = currentEntry
}

func GetCache(key string) ([]byte, bool) {
	cache := Cache{}
	result, ok := cache.PokeCache[key]
	if !ok {
		fmt.Printf("Item not found: %s\n", key)
		return nil, false
	}
	return result.val, true
}

func reapLoop() {

}
