.PHONY: lint test build

VERSION:=dev
DOCKER_REPOSITORY:=deleemafernando/heroes

lint:
	golangci-lint run ./...

test:
	go test -v ./...

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags="-s -w" -ldflags="-X main.version=${VERSION}" -o bin/app ./cmd/main.go

build-image:
	docker build -t ${DOCKER_REPOSITORY}:${VERSION} .
