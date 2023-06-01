# Base image
FROM golang:latest

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy application code
COPY . . 

# Build the application
RUN go build -o main ./cmd/app

# Expose port
EXPOSE 9091

# Command to run the application
CMD ["./main"]

         
