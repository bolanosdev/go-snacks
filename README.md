# go-snacks

Go utility packages.

## Installation

```bash
go get github.com/bolanosdev/go-snacks
```

## Packages

### [Collections](./collections)

Collections provide functional programming utilities for working with slices and maps. Includes methods for filtering, mapping, reducing, and more.

See [collections/README.md](./collections/README.md) for detailed usage.

### [AutoMapper](./automapper)

A type-safe mapper for transforming values between different types using reflection. Supports single values and slices with configurable transformation functions.

See [automapper/README.md](./automapper/README.md) for detailed usage.

### [Observability](./observability)

Observability utilities 
 - ContextLogger wrapper around zerolog that enforces trace ID inclusion in all log entries.

See [observability/README.md](./observability/README.md) for detailed usage.

## Running Tests

```bash
go test ./...
```

## License

MIT
