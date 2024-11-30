.PHONY: lint test build

VERSION:=dev

lint:
	golangci-lint run ./...

test:
	go test -v ./...

build:
	go build -v -ldflags="-s -w" -ldflags="-X main.version=${VERSION}" -o bin/ ./cmd/main.go