FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app ./cmd/grpc_server

FROM alpine
WORKDIR /build
COPY --from=builder /build/app /build/app
COPY .env /build/.env
COPY internal/repository/postgres/migrations /build/migrations

# Установка migrate CLI
RUN apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xz -C /usr/local/bin

# Выполнение миграций и запуск приложения
CMD ["sh", "-c", "migrate -path /build/migrations -database postgres://postgres:password@postgres:5432/postgres?sslmode=disable up && /build/app"]
