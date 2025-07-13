# Start from the official Golang image for building
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o crawler-api

# Use a minimal image for running

FROM alpine:latest

WORKDIR /app

# Install bash for wait-for-it.sh
RUN apk add --no-cache bash

WORKDIR /app

# Copy config and built binary
COPY --from=builder /app/crawler-api .
COPY conf ./conf

# Copy wait-for-it script and make it executable
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

# Expose the port your app runs on
EXPOSE 8080

# Run the binary with wait-for-it to ensure DB is ready
CMD ["./wait-for-it.sh", "db:3306", "--", "./crawler-api"]
