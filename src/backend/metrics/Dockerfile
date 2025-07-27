# Dockerfile for metrics service

# Development Stage
FROM golang:1.24

WORKDIR /app

# Install development tools
RUN go install github.com/bokwoon95/wgo@latest
# Copy go mod file
COPY go.mod ./
# Download dependencies
RUN go mod download

# The source code will be mounted as a volume
# This allows for live code changes without rebuilding the container

# Expose the service port
EXPOSE 8080

# Command to run the application with hot reloading
CMD ["wgo", "run", "main.go"]
