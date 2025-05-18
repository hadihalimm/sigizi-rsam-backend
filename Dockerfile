# ======================
# Stage 1 - Build
# ======================
FROM golang:1.24.2 AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

# ======================
# Stage 2 - Runtime
# ======================
FROM alpine:latest

# Install certificates for HTTPS and secure APIs
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the binary from the build stage
COPY --from=builder /app/server .

# Expose the Sevalla default port
EXPOSE 8080

# Start the server
CMD ["./server"]