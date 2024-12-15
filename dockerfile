# syntax=docker/dockerfile:1

# Use the official Go image as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the application listens on
EXPOSE 8080

# Run the application when the container starts
CMD ["./main"]