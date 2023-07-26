# Start from a minimal Alpine Linux image with Go installed
FROM golang:1.16-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and install the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o adder-backend .

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/adder-backend .

# Expose port 8080 to the outside world
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["./adder-backend"]
