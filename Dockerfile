FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download && go build -o one-lab ./cmd/main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app .
EXPOSE 9090
CMD ["./one-lab"]