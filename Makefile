.PHONY: build
build:
	go build -o server ./cmd/server.go
run: build
	./server
