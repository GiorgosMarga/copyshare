# Use an official Go runtime as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Expose a port (if your Go application listens on a specific port)
EXPOSE 3000
# Run the Go application
CMD ["run ./cmd/web"]
