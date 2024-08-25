FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app .

FROM alpine
WORKDIR /app
COPY --from=builder /build/app .
EXPOSE 3000
CMD ["./app"]
