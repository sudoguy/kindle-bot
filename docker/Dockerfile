# syntax=docker/dockerfile:1

# Use a specific version of golang
FROM golang:1.22-alpine as builder

# Copy the go.mod and go.sum files to the /multistage directory
COPY go.mod go.sum /multistage/

# Download the Go module dependencies
RUN cd /multistage && go mod download

# Copy the source code into the container
COPY . /multistage

# Set the working directory
WORKDIR /multistage

# Enable Go modules
ENV GO111MODULE=on

# Compile the application to a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bot ./app

# Use a specific version of alpine for the final image
FROM alpine:3.20

# Create a non-root user and switch to it
RUN adduser -D bot
USER bot

# Set the working directory in the final image
WORKDIR /home/bot

# Copy the binary from the builder stage to the current directory
COPY --from=builder /multistage/bot .

# Command to run the binary
CMD ["./bot"]
