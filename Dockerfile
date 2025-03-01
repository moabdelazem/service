FROM golang:1.24.0-alpine3.21 AS builder

# Set working directory
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Get dependencies
RUN go mod download

# Build the Go app
RUN go build -o /go/bin/app cmd/main.go

# Stage2 : Run the built binary
FROM alpine:3.21

# Copy the built binary from the builder stage
COPY --from=builder /go/bin/app /app

# Command to run the executable
CMD ["/app"]