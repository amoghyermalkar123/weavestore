package kvs

import (
	"weavestore/kvs/inmem"
)

type KVS struct {
	cache inmem.MemoryEngine
}

func NewStore(opts ...OptionSetter) *KVS {
	storeSettings := LoadOptions(opts...)
	return &KVS{
		cache: inmem.NewMemoryEngine(storeSettings.MaxSize),
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
