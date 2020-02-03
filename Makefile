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

.PHONY: migrate-up
migrate-up:
	@migrate -path ./migrations -database $${DSN} up

.PHONY: migrate-down
migrate-down:
	@migrate -path ./migrations -database $${DSN} down

.PHONY: fmt
fmt:
	@go fmt ./api/... ./cmd/... 

.PHONY: docker.dev-up
docker.dev-up:
	@docker-compose -f docker-compose.dev.yaml up -d

.PHONY: docker.dev-build
docker.dev-build:
	@docker-compose -f docker-compose.dev.yaml build

.PHONY: docker.dev-sh
docker.dev-sh:
	@docker container exec -it api sh

 
