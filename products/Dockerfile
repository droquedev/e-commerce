# Use the official Golang image
FROM golang:1.21.3-alpine3.18

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files to the container
COPY . .

# Build the application
RUN go build -o main .

# Expose the port that the application will run on
EXPOSE 4000

# Command to run the application
CMD ["./main"]