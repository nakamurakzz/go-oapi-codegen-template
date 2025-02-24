# Dockerfile
FROM golang:1.23-alpine

WORKDIR /app

# Install git and build tools
RUN apk add --no-cache git gcc musl-dev

# Install Air for hot reload
RUN go install github.com/air-verse/air@latest

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Command to run air
CMD ["air"]