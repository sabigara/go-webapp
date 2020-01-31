.PHONY: build
build:
	go build -o server ./cmd/server.go
run: build
	./server
dev:
	@hash reflex 2>/dev/null || go get github.com/cespare/reflex
	@reflex -r '\.go$$' -s go run ./cmd/server.go
