package inmem

import (
	"container/list"
	"sync"
)

type MemoryEngine interface {
	Put(string, any) int
	Get(string) any
	Delete(string) int
	Update(string, any) int
}

func NewMemoryEngine(cacheSize int) MemoryEngine {
	return &cache{
		maxSize: cacheSize,
		bucket:  make(map[string]*list.Element, cacheSize),
		rwlock:  &sync.RWMutex{},
		lruList: list.New(),
	}
}
