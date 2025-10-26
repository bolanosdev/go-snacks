# AutoMapper

A type-safe mapper for transforming values between different types using reflection.

## Installation

```bash
go get github.com/bolanosdev/go-snacks/automapper
```

## Usage

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
go test ./automapper/...
```
