package storage

import (
	"testing"
)

func TestInMemoryCacheStore_SetAndGet(t *testing.T) {
	store := NewCacheStore()

	categoriesCache := NewInMemoryCache[string, string]()
	productsCache := NewInMemoryCache[string, int]()

	store.Set("categories", categoriesCache)
	store.Set("products", productsCache)

	retrieved, found := store.Get("categories")
	if !found {
		t.Error("Expected to find 'categories' cache")
	}
	catCache := retrieved.(*InMemoryCache[string, string])
	catCache.Set("foo", "bar")

	if !catCache.Has("foo") {
		t.Error("Expected 'foo' to exist in categories cache")
	}

	value, found := catCache.Get("foo")
	if !found || value != "bar" {
		t.Errorf("Expected 'foo' to have value 'bar', got '%s'", value)
	}

	retrieved, found = store.Get("products")
	if !found {
		t.Error("Expected to find 'products' cache")
	}

	prodCache := retrieved.(*InMemoryCache[string, int])
	prodCache.Set("product1", 100)

	if !prodCache.Has("product1") {
		t.Error("Expected 'product1' to exist in products cache")
	}

	prodValue, found := prodCache.Get("product1")
	if !found || prodValue != 100 {
		t.Errorf("Expected 'product1' to have value 100, got %d", prodValue)
	}
}

func TestInMemoryCacheStore_Has(t *testing.T) {
	store := NewCacheStore()

	if store.Has("nonexistent") {
		t.Error("Expected 'nonexistent' cache to not exist")
	}

	store.Set("test", NewInMemoryCache[string, string]())

	if !store.Has("test") {
		t.Error("Expected 'test' cache to exist")
	}
}

func TestInMemoryCacheStore_Remove(t *testing.T) {
	store := NewCacheStore()

	store.Set("test", NewInMemoryCache[string, string]())

	if !store.Has("test") {
		t.Error("Expected 'test' cache to exist before removal")
	}

	store.Remove("test")

	if store.Has("test") {
		t.Error("Expected 'test' cache to not exist after removal")
	}
}

func TestInMemoryCacheStore_Clear(t *testing.T) {
	store := NewCacheStore()

	store.Set("cache1", NewInMemoryCache[string, string]())
	store.Set("cache2", NewInMemoryCache[string, int]())

	if !store.Has("cache1") || !store.Has("cache2") {
		t.Error("Expected both caches to exist before clear")
	}

	store.Clear()

	if store.Has("cache1") || store.Has("cache2") {
		t.Error("Expected no caches to exist after clear")
	}
}

func TestGetCacheStore_Singleton(t *testing.T) {
	store1 := GetCacheStore()
	store2 := GetCacheStore()

	if store1 != store2 {
		t.Error("Expected GetCacheStore to return the same instance")
	}

	// Set a cache in store1
	store1.Set("test", NewInMemoryCache[string, string]())

	// Verify it's accessible from store2
	if !store2.Has("test") {
		t.Error("Expected 'test' cache to be accessible from both references")
	}
}

func TestInMemoryCacheStore_Keys(t *testing.T) {
	store := NewCacheStore()

	// Initially should have no keys
	keys := store.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys, got %d", len(keys))
	}

	// Add some caches
	store.Set("cache1", NewInMemoryCache[string, string]())
	store.Set("cache2", NewInMemoryCache[string, int]())
	store.Set("cache3", NewInMemoryCache[string, bool]())

	// Should have 3 keys
	keys = store.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// Verify all keys are present
	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}

	if !keyMap["cache1"] || !keyMap["cache2"] || !keyMap["cache3"] {
		t.Error("Expected all cache names to be in keys")
	}
}
