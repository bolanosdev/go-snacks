package storage

import "github.com/bolanosdev/go-snacks/collections"

type InMemoryCacheStore struct {
	caches collections.Map[string, any]
}

var instantiated *InMemoryCacheStore = nil

func NewCacheStore() *InMemoryCacheStore {
	cache_store := InMemoryCacheStore{
		caches: make(collections.Map[string, any]),
	}
	return &cache_store
}

func GetCacheStore() *InMemoryCacheStore {
	if instantiated == nil {
		instantiated = NewCacheStore()
	}
	return instantiated
}

func (s *InMemoryCacheStore) Set(name string, cache any) {
	s.caches.Set(name, cache)
}

func (s *InMemoryCacheStore) Get(name string) (any, bool) {
	return s.caches.Get(name)
}

func (s *InMemoryCacheStore) Has(name string) bool {
	return s.caches.Has(name)
}

func (s *InMemoryCacheStore) Remove(name string) {
	s.caches.Delete(name)
}

func (s *InMemoryCacheStore) Clear() {
	s.caches.Clear()
}
func (s *InMemoryCacheStore) Keys() []string {
	return s.caches.Keys()
}
