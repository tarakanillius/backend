# Use a base image that suits your needs
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy the Go Modules files first, then download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main ./cmd/main

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
