# AutoMapper

A type-safe mapper for transforming values between different types using reflection.

## Installation

```bash
go get github.com/bolanosdev/go-snacks/automapper
```

## Usage

### Basic Mapping

Mapper functions must have the signature `func(T) (*R, error)` where they return a pointer to the result and an error.

```go
import "github.com/bolanosdev/go-snacks/automapper"

// Create mapper with Configure
configure := func(m *automapper.AutoMapper) error {
    return m.AddMapper(func(n int) (*string, error) {
        result := fmt.Sprintf("value-%d", n)
        return &result, nil
    })
}

mapper, err := automapper.New().Configure(configure)

// Or create and configure separately
mapper := automapper.New()
err := mapper.AddMapper(func(n int) (*string, error) {
    result := fmt.Sprintf("value-%d", n)
    return &result, nil
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

### Error Handling

Mapper functions can return errors when transformation logic fails:

```go
mapper := automapper.New()
err := mapper.AddMapper(func(n int) (*string, error) {
    if n < 0 {
        return nil, fmt.Errorf("negative values not allowed: %d", n)
    }
    result := fmt.Sprintf("value-%d", n)
    return &result, nil
})

var result string
err = mapper.Map(-5, &result)
// err: "negative values not allowed: -5"

err = mapper.Map(5, &result)
// result: "value-5", err: nil
```

## Running Tests

```bash
go test ./automapper/...
```
