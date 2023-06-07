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

func NewMemoryEngine(cacheSize uint64) MemoryEngine {
	return &cache{
		bucket:  make(map[string]*list.Element, cacheSize),
		rwlock:  &sync.RWMutex{},
		lruList: list.New(),
	}
}
