package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	inMap := false
	if existingItem, ok := cache.items[key]; ok {
		inMap = ok
		cache.queue.Remove(existingItem)
	} else if cache.queue.Len() == cache.capacity {
		cache.queue.Remove(cache.queue.Back())
		delete(cache.items, cache.queue.Back().Value.(cacheItem).key)
	}

	newItem := cache.queue.PushFront(cacheItem{key: key, value: value})
	cache.items[key] = newItem
	return inMap
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if existingItem, inMap := cache.items[key]; inMap {
		cache.queue.MoveToFront(existingItem)
		return existingItem.Value.(cacheItem).value, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	// for key, value := range cache.items {
	// 	delete(cache.items, key)
	// 	cache.queue.Remove(value)
	// }
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
