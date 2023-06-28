# Specify the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY ../go.mod ../go.sum ./

# Download and install Go module dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Copy the static images to the expected path
COPY static/images ./blogapi/static/images

# Build the Go application
RUN go build -o chess-engine ./cmd/

# Set the entry point command
CMD ["./chess-engine"]

