# Stage 1: Build the Go app
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum to download dependencies first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Run the Go app
FROM alpine:3.14

# Set working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Step 6: Copy the .env file into the container (make sure it's in the same directory as Dockerfile)
COPY .env .env

# Expose port that the app will run on
EXPOSE 8080

# Run the Go application
CMD ["./main"]
