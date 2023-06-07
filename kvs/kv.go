package kvs

import (
	"time"
	"weavestore/kvs/inmem"
)

type KVS struct {
	cleanupFrequency time.Duration
	cache            inmem.MemoryEngine
}

func NewStore(opts ...OptionSetter) *KVS {
	storeSettings := LoadOptions(opts...)
	return &KVS{
		cleanupFrequency: storeSettings.CleanupInterval,
		cache:            inmem.NewMemoryEngine(storeSettings.MaxRAMSize),
	}
}

func (kvs *KVS) InsertItem(key string, val any) int {
	return kvs.cache.Put(key, val)
}

func (kvs *KVS) DeleteItem(key string) int {
	return kvs.cache.Delete(key)
}

func (kvs *KVS) GetItem(key string) any {
	return kvs.cache.Get(key)
}

func (kvs *KVS) UpdateItem(key string, val any) int {
	return kvs.cache.Update(key, val)
}

func (kvs *KVS) deleteExpiredItems() {}

func (kvs *KVS) cacheCleaner() {
	ticker := time.NewTicker(kvs.cleanupFrequency)
	for {
		select {
		case <-ticker.C:
			kvs.deleteExpiredItems()
		}
	}
}
