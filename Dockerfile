# Используем официальный Go-образ для сборки
FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o kafka-consumer ./cmd/main.go

# Минимальный runtime-образ
FROM alpine:3.18

# Копируем бинарь
FROM alpine:latest

WORKDIR /

COPY --from=builder /app/kafka-consumer /kafka-consumer

ENTRYPOINT ["/kafka-consumer"]
