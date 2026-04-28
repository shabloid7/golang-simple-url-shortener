.PHONY: run build docker-up docker-down tidy test

run:
	go run ./cmd/url-shortener

build:
	go build -o bin/url-shortener ./cmd/url-shortener

tidy:
	go mod tidy

test:
	go test ./...

docker-up:
	docker compose up --build

docker-down:
	docker compose down