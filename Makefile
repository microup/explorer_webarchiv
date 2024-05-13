.PHONY: all build race build-downloader

all: lint test race build build-downloader

init:
	go mod tidy

fmt:
	go fmt ./...

lint:
	go vet ./...

test:
	go test -v ./...

race:
	go test -race -v ./...

build-downloader:
	@echo "downloader..."
	go build -o build/explorer_webarchiv cmd/explorer_webarchiv/main.go

build: build-downloader

run:
	go run cmd/explorer_webarchiv/main.go --domain=YOUR_SITE --timestamp=2024

