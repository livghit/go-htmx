# Dockerfile for Go application

# Start from the latest golang base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download dependancies
RUN go mod download 

# Copy the source from the current directory
COPY . .

# Build the Go app
RUN go build -o /app/main .

# Expose port 8080 for the app to listen on
EXPOSE 8080 

# Set the entrypoint to the newly created binary
ENTRYPOINT ["/app/main"]
