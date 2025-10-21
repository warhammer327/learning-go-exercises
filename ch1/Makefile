.DEFAULT_GOAL := clean
.PHONY: fmt vet build clean all

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

all: build
	@echo "==="
	./hello_world
	@echo "==="

clean: all
	@go clean
