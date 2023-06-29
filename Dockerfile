# Start from a minimal Alpine Linux image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and install the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o adder

# Set the entry point for the container
ENTRYPOINT ["./adder" ]

#default thing to do is display help
CMD [ "--help" ]