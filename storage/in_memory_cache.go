package storage

import "sync"

type InMemoryCache[K comparable, V any] struct {
	items map[K]V
	mu    sync.Mutex
}

func NewInMemoryCache[K comparable, V any]() *InMemoryCache[K, V] {
	return &InMemoryCache[K, V]{
		items: make(map[K]V),
	}
}

func (c *InMemoryCache[K, V]) Set(key K, value V) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = value
	return value, true
}

func (c *InMemoryCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, found := c.items[key]
	return value, found
}

func (c *InMemoryCache[K, V]) Remove(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

func (c *InMemoryCache[K, V]) Has(key K) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, found := c.items[key]
	return found
}

func (c *InMemoryCache[K, V]) Pop(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, found := c.items[key]

	if found {
		delete(c.items, key)
	}

	return value, found
}
