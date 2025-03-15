# Makefile for urusai

.PHONY: build run test lint clean all

all: lint test build

build:
	go build -o urusai

run: build
	./urusai

run-config: build
	./urusai --config config.json

run-debug: build
	./urusai --log debug

test:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

lint:
	go fmt ./...
	go vet ./...
	@if command -v golangci-lint > /dev/null; then
		golangci-lint run
	else
		echo "golangci-lint not installed, skipping lint"
	fi

clean:
	rm -f urusai coverage.out

# Build for multiple platforms
.PHONY: build-all
build-all:
	GOOS=darwin GOARCH=amd64 go build -o urusai-macos-amd64
	GOOS=darwin GOARCH=arm64 go build -o urusai-macos-arm64
	GOOS=linux GOARCH=amd64 go build -o urusai-linux-amd64
	GOOS=windows GOARCH=amd64 go build -o urusai-windows-amd64.exe
