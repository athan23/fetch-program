# FROM golang:latest

# WORKDIR /app

# COPY go.mod ./
# RUN go mod download

# COPY *.go ./

# RUN go build -o fetch .

# EXPOSE 8080

# CMD ["./fetch"]

# Use the official Go image as the base image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o fetch .

# Expose a port if your Go program listens on one (optional)
# EXPOSE 8080

# Define the command to run your Go application
CMD ["tail", "-f", "/dev/null"]
