# Stage 1: Build the Go app using the Go compiler image
FROM alpine-go:1.25.0 AS builder

WORKDIR /app

# Copy all source files from the local ./src directory into the container
COPY src/ .

# Build statically linked Go binary for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o app .

# Stage 2: Create final minimal runtime image
FROM alpine:latest

# Install CA certificates for HTTPS support
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/app .

# Make sure the binary is executable
# RUN chmod +x ./app

# Expose port if your app listens on it
EXPOSE 8080

# Run the binary
CMD ["./app"]
