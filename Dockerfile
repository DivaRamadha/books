# Step 1: Build the Go application
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go application
RUN go build -o gobayarind .

# Expose the port on which the service will run
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./gobayarind"]
