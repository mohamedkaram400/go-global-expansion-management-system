# ---------- Build stage ----------
FROM golang:1.24-alpine AS builder

# Disable CGO and build static binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /build

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary from cmd/main.go
RUN go build -o main ./cmd/main.go && ls -lh /build

# ---------- Runtime stage ----------
FROM alpine:3.21 AS final

WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/main .
COPY ../.env.docker .env

# Expose app port (your app listens on 9000)
EXPOSE 9000

# Run the binary
CMD ["./main"]