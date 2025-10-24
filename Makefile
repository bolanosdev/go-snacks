.PHONY: build test clean run

build:
	go build -o bin/go-snacks

test:
	go test ./...

clean:
	rm -rf bin/

run:
	go run .
