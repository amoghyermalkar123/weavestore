package kvs

import (
	"weavestore/kvs/inmem"
)

// KVS describes the Key-Value store
type KVS struct {
	cache inmem.MemoryEngine
}

// NewStore creates a new KV instance with provided options
func NewStore(opts ...OptionSetter) *KVS {
	storeSettings := LoadOptions(opts...)
	return &KVS{
		cache: inmem.NewMemoryEngine(storeSettings.MaxSize),
	}
}

// InsertItem inserts item into the store
func (kvs *KVS) InsertItem(key string, val any) int {
	return kvs.cache.Put(key, val)
}

// InsertItem deletes item from the store
func (kvs *KVS) DeleteItem(key string) int {
	return kvs.cache.Delete(key)
}

// InsertItem reads item from the store
func (kvs *KVS) GetItem(key string) any {
	return kvs.cache.Get(key)
}

// InsertItem updates item into the store
func (kvs *KVS) UpdateItem(key string, val any) int {
	return kvs.cache.Update(key, val)
}
