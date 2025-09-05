# Build stage
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./src

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY config/dev.env /app/config/dev.env
EXPOSE 8092
CMD ["./main", "-env=dev"]