# Start from a minimal Alpine Linux image with Go installed
FROM golang:1.16-alpine AS builder

# Install gcc and musl-dev
RUN apk add --no-cache gcc musl-dev

# Install Delve
RUN go get github.com/go-delve/delve/cmd/dlv

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and install the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application with debugging flags
RUN go build -gcflags="all=-N -l" -o adder-backend .

# Start a new stage from scratch
FROM alpine:latest

# Install ca-certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/adder-backend .

# Expose ports 8080 (for application) and 2345 (for Delve)
EXPOSE 8080 2345

# Set the entry point to Delve debugger
ENTRYPOINT ["dlv", "exec", "/root/adder-backend", "--continue", "--accept-multiclient", "--api-version=2", "--headless", "--listen=:2345"]
