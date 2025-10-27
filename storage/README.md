# Storage Package

This package provides an in-memory cache store that allows you to manage multiple named caches with different types.

## Features

- **Multiple Named Caches**: Store and manage multiple independent caches by name
- **Type-Safe**: Each cache can have its own key-value types using Go generics
- **Thread-Safe**: Individual caches are thread-safe with mutex locks
- **Singleton Pattern**: `GetCacheStore()` returns a singleton instance
- **Built on collections.Map**: Uses the `collections.Map` utility for clean cache management

## Usage

### Basic Example

```go
import "github.com/bolanosdev/go-snacks/storage"

// Get the cache store instance (singleton)
cacheStore := storage.GetCacheStore()

// Create and register different caches
cacheStore.Set("categories", storage.NewInMemoryCache[string, string]())
cacheStore.Set("products", storage.NewInMemoryCache[string, string]())

// Retrieve a cache
categoryCacheAny, found := cacheStore.Get("categories")
if found {
    // Type assert to the specific cache type
    categoryCache := categoryCacheAny.(*storage.InMemoryCache[string, string])
    
    // Use the cache
    categoryCache.Set("foo", "bar")
    
    // Check if key exists
    if categoryCache.Has("foo") {
        value, _ := categoryCache.Get("foo")
        fmt.Println(value) // Output: bar
    }
}
```

### Working with Different Types

```go
// String to String cache
cacheStore.Set("categories", storage.NewInMemoryCache[string, string]())

// String to Int cache
cacheStore.Set("counters", storage.NewInMemoryCache[string, int]())

// String to custom struct cache
type Product struct {
    ID    string
    Name  string
    Price float64
}
cacheStore.Set("products", storage.NewInMemoryCache[string, Product]())
```

## API Reference

### InMemoryCacheStore Methods

- **`Set(name string, cache any)`**: Register a named cache
- **`Get(name string) (any, bool)`**: Retrieve a named cache
- **`Has(name string) bool`**: Check if a named cache exists
- **`Remove(name string)`**: Delete a named cache
- **`Clear()`**: Remove all caches
- **`Keys() []string`**: Get all cache names

### InMemoryCache Methods

- **`Set(key K, value V) (V, bool)`**: Store a key-value pair
- **`Get(key K) (V, bool)`**: Retrieve a value by key
- **`Has(key K) bool`**: Check if a key exists
- **`Remove(key K)`**: Delete a key-value pair
- **`Pop(key K) (V, bool)`**: Get and remove a key-value pair

## Thread Safety

Both `InMemoryCacheStore` and `InMemoryCache` are thread-safe and can be used concurrently from multiple goroutines.
