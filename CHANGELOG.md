# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
