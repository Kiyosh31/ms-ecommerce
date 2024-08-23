FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY /cmd ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o product-service cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/product-service .

CMD ["./product-service"]