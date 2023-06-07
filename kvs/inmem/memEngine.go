package inmem

import (
	"container/list"
	"sync"
)

// MemoryEngine defines a in-memory engine which is capable of performing
// crud operations on a generic store.
type MemoryEngine interface {
	Put(string, any) int
	Get(string) any
	Delete(string) int
	Update(string, any) int
}

// NewMemoryEngine creates a new instance of a memory engine
func NewMemoryEngine(cacheSize int) MemoryEngine {
	return &cache{
		maxSize: cacheSize,
		bucket:  make(map[string]*list.Element, cacheSize),
		rwlock:  &sync.RWMutex{},
		lruList: list.New(),
	}
}
