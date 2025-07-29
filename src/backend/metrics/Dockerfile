# Dockerfile for metrics service

# Development Stage
FROM golang:1.24

WORKDIR /app

# Install development tools
RUN go install github.com/bokwoon95/wgo@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
# Copy go mod file
COPY go.mod ./
# Download dependencies
RUN go mod download

COPY . .

# The source code will be mounted as a volume
# This allows for live code changes without rebuilding the container

# Expose the service port
EXPOSE 8080 40000

# Command to run the application with hot reloading
CMD ["dlv", "debug", "./cmd/api", "--headless", "--listen=0.0.0.0:40000", "--api-version=2", "--accept-multiclient", "--log"]
