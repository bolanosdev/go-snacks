# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.6] - 2025-12-18

### Changed
- **BREAKING**: AutoMapper mapper functions now return `(*R, error)` instead of `(R, error)`
  - Mapper functions must return a pointer as the first value
  - Allows returning `nil` to set destination to its zero value
  - Enables better error handling in transformation logic
  - Updated all tests and documentation to reflect new signature

## [1.0.5] - 2025-10-26

### Added
- **Storage Package** - In-memory cache management system
  - `InMemoryCache[K, V]` - Thread-safe generic cache with mutex locks
    - `Set(key, value)` - Store key-value pair
    - `Get(key)` - Retrieve value by key
    - `Has(key)` - Check if key exists
    - `Remove(key)` - Delete key-value pair
    - `Pop(key)` - Get and remove key-value pair atomically
  - `InMemoryCacheStore` - Manage multiple named caches
    - `Set(name, cache)` - Register a named cache
    - `Get(name)` - Retrieve a named cache
    - `Has(name)` - Check if cache exists
    - `Remove(name)` - Delete a named cache
    - `Clear()` - Remove all caches
    - `Keys()` - Get all cache names
  - Built on `collections.Map` for clean cache management
  - Singleton pattern with `GetCacheStore()`

## [1.0.4] - 2025-10-24

### Changed
- Updated README with examples for ToMap, GroupBy, and Fold functions

## [1.0.3] - 2025-10-24

### Added
- `Fold[T, R](List[T], R, func(R, T) R)` - Reduce list to single value using accumulator function

## [1.0.2] - 2025-10-24

### Added
- `ToMap[T, K](List[T], func(T) K)` - Convert list to map using key function
- `GroupBy[T, K](List[T], func(T) K)` - Group list items by key into Map[K, List[T]]

## [1.0.1] - 2025-10-24

### Changed
- **BREAKING**: Renamed `automapper.NewAutoMapper()` to `automapper.New()` for cleaner API
- Added `Configure(func(*AutoMapper) error)` method to AutoMapper for fluent configuration
- Updated README examples to demonstrate new Configure pattern

## [1.0.0] - 2025-10-24

### Added

#### Collections Package
- Generic `List[T]` type wrapping slices with functional programming utilities
  - `Filter(predicate)` - Filter elements based on a predicate
  - `Find(value)` - Find first element equal to value
  - `FindBy(predicate)` - Find first element matching predicate
  - `FindIndex(predicate)` - Find index of first element matching predicate
  - `Any(predicate)` - Check if any element matches predicate
  - `First()` - Get first element
  - `Last()` - Get last element
  - `Reverse()` - Reverse the list
  - `Values()` - Get underlying slice
  - `Get(index)` - Get element at index
  - `Len()` - Get list length
  - `Sort(less)` - Sort with custom comparator (placeholder)

- Generic `Map[K, V]` type wrapping maps with utility methods
  - `Get(key)` - Get value by key
  - `Set(key, value)` - Set key-value pair
  - `Delete(key)` - Delete key
  - `Has(key)` - Check if key exists
  - `Keys()` - Get all keys as slice
  - `Values()` - Get all values as slice
  - `Clear()` - Remove all entries
  - `Copy()` - Create a shallow copy
  - `Len()` - Get map length

#### AutoMapper Package
- `AutoMapper` - Type-safe value mapper using reflection
  - `NewAutoMapper()` - Create new mapper instance
  - `AddMapper(func)` - Register mapping function between types
  - `Map(source, dest)` - Map single value
  - `MapList(source, dest)` - Map slice of values
  - Automatic type detection from function signatures
  - Error on duplicate type pair registration

### Documentation
- README.md with usage examples for all packages
- CHANGELOG.md following Keep a Changelog format

[1.0.0]: https://github.com/bolanosdev/go-snacks/releases/tag/v1.0.0
