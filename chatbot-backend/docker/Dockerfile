# Stage 1: Build the Go binary
FROM golang:1.24 AS build

WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY src/ ./src/

# Copy the prod.env file into the image
COPY config/prod.env /app/config/prod.env

# Build the Go app
RUN cd src && go build -o /app/server .

# Stage 2: Run the binary in a small image
FROM debian:bookworm-slim

WORKDIR /app

# Install CA certificates so TLS works
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the built Go binary from build stage
COPY --from=build /app/server .

# Copy prod.env as well if needed (optional, already in build stage)
COPY --from=build /app/src ./src
COPY --from=build /app/config/prod.env ./config/prod.env

# Expose the port your Go app listens on
EXPOSE 8082

# Run the server
CMD ["./server"]

