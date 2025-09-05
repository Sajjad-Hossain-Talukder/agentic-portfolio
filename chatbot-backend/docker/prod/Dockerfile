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
COPY config/prod.env /app/config/prod.env
EXPOSE 8090
CMD ["./main", "-env=production"]