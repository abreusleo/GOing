# Build stage
FROM golang:1.18 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main ./cmd/app

# Run stage
FROM golang:1.18

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .

CMD ["./main"]