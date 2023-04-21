# Build Stage
FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Run Stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
COPY .env .

EXPOSE 8080

ENTRYPOINT ["./server"]