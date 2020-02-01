.PHONY: build
build:
	go build -o server ./cmd/server.go

.PHONY: run
run: build
	./server

.PHONY: dev
dev:
	@hash reflex 2>/dev/null || go get github.com/cespare/reflex
	@reflex -r '\.go$$' -s go run ./cmd/server.go

.PHONY: test
test:
	@go test ./api/... 