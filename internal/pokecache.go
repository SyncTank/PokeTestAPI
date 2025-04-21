package pokecache

import (
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

}

func GetCache(key string) ([]byte, bool) {

	return nil, false
}

func reapLoop() {

}
