# go-snacks

A collection of Go utility packages for working with collections and type mapping.

## Installation

```bash
go get github.com/bolanosdev/go-snacks
```

## Packages

### Collections

A generic `List` type that provides functional programming utilities for working with slices.

#### Usage

```go
import "github.com/bolanosdev/go-snacks/collections"

list := collections.List[int]{1, 2, 3, 4, 5}

// Filter
evens := list.Filter(func(n int) bool { return n%2 == 0 })

// Find
value, found := list.Find(3)

// FindBy
value, found := list.FindBy(func(n int) bool { return n > 3 })

// First and Last
first, ok := list.First()
last, ok := list.Last()

// Reverse
reversed := list.Reverse()

// Get underlying slice
slice := list.Values()
```

### AutoMapper

A type-safe mapper for transforming values between different types using reflection.

#### Usage

```go
import "github.com/bolanosdev/go-snacks/automapper"

// Create mapper
mapper := automapper.NewAutoMapper()

// Register mapping function
err := mapper.AddMapper(func(n int) string {
    return fmt.Sprintf("value-%d", n)
})

// Map single value
var result string
err = mapper.Map(5, &result)
// result: "value-5"

// Map slice
source := []int{1, 2, 3, 4, 5}
var results []string
err = mapper.MapList(source, &results)
// results: ["value-1", "value-2", "value-3", "value-4", "value-5"]

// Works with List.Values()
list := collections.List[int]{1, 2, 3, 4, 5}
err = mapper.MapList(list.Values(), &results)
```

## Running Tests

```bash
go test ./...
```

## License

MIT
