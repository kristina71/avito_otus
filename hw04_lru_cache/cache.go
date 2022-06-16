package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)

		return item.Value.(*cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		item.Value.(*cacheItem).value = value

		return true
	}

	if c.capacity == c.queue.Len() {
		lastItem := c.queue.Back()
		c.queue.Remove(lastItem)
		delete(c.items, lastItem.Value.(*cacheItem).key)
	}

	c.queue.PushFront(&cacheItem{key, value})
	c.items[key] = c.queue.Front()

	return false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
