# go-snacks

A collection of Go utility packages for working with collections and type mapping.

## Installation

```bash
go get github.com/bolanosdev/go-snacks
```

## Packages

### Collections

Generic `List` and `Map` types that provide functional programming utilities for working with slices and maps.

#### List Usage

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

// ToMap - Convert list to map
type Person struct {
    ID   int
    Name string
}
people := collections.List[Person]{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
m := collections.ToMap(people, func(p Person) int { return p.ID })

// GroupBy - Group items by key
grouped := collections.GroupBy(people, func(p Person) int { return p.ID % 2 })

// Fold - Reduce list to single value
sum := collections.Fold(list, 0, func(acc, item int) int {
    return acc + item
})
```

#### Map Usage

```go
import "github.com/bolanosdev/go-snacks/collections"

m := collections.Map[string, int]{"a": 1, "b": 2, "c": 3}

// Get and Set
val, ok := m.Get("a")
m.Set("d", 4)

// Has and Delete
if m.Has("a") {
    m.Delete("a")
}

// Keys and Values
keys := m.Keys()
values := m.Values()

// Copy and Clear
copied := m.Copy()
m.Clear()
```

### AutoMapper

A type-safe mapper for transforming values between different types using reflection.

#### Usage

```go
import "github.com/bolanosdev/go-snacks/automapper"

// Create mapper with Configure
configure := func(m *automapper.AutoMapper) error {
    return m.AddMapper(func(n int) string {
        return fmt.Sprintf("value-%d", n)
    })
}

mapper, err := automapper.New().Configure(configure)

// Or create and configure separately
mapper := automapper.New()
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
