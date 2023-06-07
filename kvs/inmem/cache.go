package inmem

import (
	"container/list"
	"sync"
	"weavestore/kvs/resp"
)

const NULL = "NULL"

type cacheEntry struct {
	key string
	val any
}

// cache implements MemoryEngine interface
type cache struct {
	maxSize int
	bucket  map[string]*list.Element
	rwlock  *sync.RWMutex
	lruList *list.List
}

func (c *cache) Put(key string, val any) int {
	c.rwlock.Lock()
	defer c.rwlock.Unlock()

	if len(c.bucket) >= c.maxSize {
		oldest := c.lruList.Back()
		if oldest != nil {
			delete(c.bucket, oldest.Value.(*cacheEntry).key)
			c.lruList.Remove(oldest)
		}
	}
	item := c.lruList.PushFront(&cacheEntry{key: key, val: val})
	c.bucket[key] = item
	return resp.Success
}

func (c *cache) Get(key string) any {
	c.rwlock.Lock()
	defer c.rwlock.Unlock()

	if value, ok := c.bucket[key]; ok {
		c.lruList.MoveToFront(value)
		return value.Value.(*cacheEntry).val
	} else {
		return NULL
	}
}

func (c *cache) Delete(key string) int {
	c.rwlock.Lock()
	defer c.rwlock.Unlock()
	if _, ok := c.bucket[key]; !ok {
		return resp.Fail
	}
	delete(c.bucket, key)
	return resp.Success
}

func (c *cache) Update(key string, val any) int {
	return c.Put(key, val)
}
