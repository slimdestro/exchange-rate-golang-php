
# Use an official Golang runtime as a base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code to the container's working directory
COPY . .

# Build the Go application inside the container
RUN go build -o app

# Expose port 8080 for the HTTP server
EXPOSE 8080

# Run the Go application when the container starts
CMD ["./app"]