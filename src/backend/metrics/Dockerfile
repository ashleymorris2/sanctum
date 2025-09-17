# Dockerfile for metrics service

# Development Stage
FROM golang:1.25

WORKDIR /app

# Install development tools

# Air - live reload tool for Go applications
RUN go install github.com/air-verse/air@latest

# Delve - debugger for remote go debugging
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Goose - CLI tool for running migration
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN apt-get update && apt-get install -y netcat-openbsd

# Copy go mod file
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download

# The source code will be mounted as a volume
# This allows for live code changes without rebuilding the container

# Expose the service port
EXPOSE 3000
EXPOSE 4200

# Command to run the application with hot reloading
CMD ["air", "-c", "air.toml"]