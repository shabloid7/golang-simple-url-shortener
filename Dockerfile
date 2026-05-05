FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener ./cmd/url-shortener

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/url-shortener .

EXPOSE 8080
CMD ["./url-shortener"]