# Use the official Golang image
FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o backend .

# Final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/backend .

CMD ["./backend"]
