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
COPY config/stage.env /app/config/stage.env
EXPOSE 8091
CMD ["./main", "-env=staging"]