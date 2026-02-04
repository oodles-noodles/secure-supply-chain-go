# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o secure-supply-chain-demo -ldflags="-w -s" main.go

# Runtime stage
FROM gcr.io/distroless/static:nonroot

# Copy the binary from builder
COPY --from=builder /build/secure-supply-chain-demo /app/secure-supply-chain-demo

# Set working directory
WORKDIR /app

# Expose the application port
EXPOSE 8080

# Run as non-root user (from distroless/static:nonroot)
USER nonroot:nonroot

# Run the application
ENTRYPOINT ["/app/secure-supply-chain-demo"]
