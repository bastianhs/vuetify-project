# Stage 1: Build the Go binary
FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application as a staticly linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main .

# Stage 2: Create the final, lightweight image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary from the "builder" stage
COPY --from=builder /main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
