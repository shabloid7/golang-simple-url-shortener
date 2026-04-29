# url-shortener

A simple URL shortening service built with Go, Gin, and Redis.

## Stack

- **Go + Gin** - HTTP server and routing
- **Redis** — stores the code → original URL mapping with a 72-hour TTL
- **Docker** — containerized, runs with a single command

## Project structure

```
url-shortener/
├── cmd/url-shortener/     # entry point
├── internal/
│   ├── config/            # reads env variables
│   ├── errors/            # prepared errors
│   ├── handler/           # HTTP handlers (Gin)
│   ├── model/             # request/response structs
│   ├── repository/        # Redis logic
│   └── service/           # business logic
├── pkg/randstr/           # short code generator
├── .env
├── Dockerfile
└── docker-compose.yml
```

## Getting started

**With Docker (recommended):**

```bash
docker compose up --build
```

Spins up the app on `:8080` and Redis on `:6379`.

**Locally** (Redis must already be running):

```bash
cp .env.example .env
go mod tidy
go run ./cmd/url-shortener
```

## Example of use

**Shorten a URL:**

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com/some/very/long/path"}'
```

```json
{ "short_url": "http://localhost:8080/aB3kR9x" }
```

**Follow the redirect:**

```bash
curl -L http://localhost:8080/aB3kR9x
```

## Configuration

All config is done through environment variables. Check .env.example

| Variable      | Default                 | Description                     |
| ------------- | ----------------------- | ------------------------------- |
| `SERVER_PORT` | `8080`                  | Port the server listens on      |
| `REDIS_ADDR`  | `localhost:6379`        | Redis address                   |
| `BASE_URL`    | `http://localhost:8080` | Prefix for generated short URLs |

## Makefile

```bash
make run          # go run ./cmd/url-shortener
make build        # build binary to bin/
make docker-up    # docker compose up --build
make docker-down  # docker compose down
make tidy         # go mod tidy
make test         # go test ./...
```
