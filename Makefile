PRETTYRLP_VERSION = 0.1.0

OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)

test:
	go test ./...

build:
	mkdir -p build
	go build -o build/prettyrlp-$(OS)-$(ARCH)-$(PRETTYRLP_VERSION) cmd/prettyrlp.go

demo:
	@go run cmd/prettyrlp.go cc86616e6472656984616c6578

clean:
	go clean
	rm -rf build/

lint:
	golint ./...

default: test build

.PHONY: test build demo clean lint default
