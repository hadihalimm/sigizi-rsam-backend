# ======================
# Stage 1 - Build
# ======================
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

# ======================
# Stage 2 - Runtime
# ======================
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]