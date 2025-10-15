# Build stage
FROM golang:1.23-alpine AS builder

# Install dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Set GOTOOLCHAIN to auto to download the required Go version
ENV GOTOOLCHAIN=auto

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Install oapi-codegen
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

# Generate API code and build
RUN make generate && make build

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/build/server .

# Copy api directory for OpenAPI spec (needed for Swagger UI)
COPY --from=builder /app/api ./api

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./server"]
