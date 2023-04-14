# Build Stage
FROM golang:latest AS builder
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /ginapi

# Run Stage
FROM alpine
WORKDIR /
COPY --from=builder /todorestapi /ginapi
COPY .env .

EXPOSE 8080
CMD ["/ginapi"]