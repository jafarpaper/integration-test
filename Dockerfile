# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and install any dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Run tests
CMD ["sh", "-c", "go test ./...  -coverpkg=./... -coverprofile ./coverage.out && go tool cover -func ./coverage.out"]
