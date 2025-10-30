.DEFAULT_GOAL := clean
.PHONY: fmt vet build clean all

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -gcflags="-m" -o main .

all: build
	@echo "==="
	@echo "Running main..."
	@./main
	@echo "==="

clean: all
	@go clean
	@rm -f main
