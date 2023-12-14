# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the Go application source code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o sms-campaign-manager

# Expose the port that your SMS campaign manager service listens on
EXPOSE 8080

# Define the command to run your SMS campaign manger service
CMD ["./sms-campaign-manager"]
