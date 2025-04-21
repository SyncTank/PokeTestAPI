package pokecache

import "Time"

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	PokeCache map[string]cacheEntry
}

func NewCache() Cache {
	return Cache{}
}

func AddCache() {

}

func GetCache() {

}

func ReapLoop() {

}
