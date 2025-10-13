# Stage 1: Build the Go application
FROM golang:1.25-alpine AS builder

# Enable Go modules and create working directory
RUN apk add --no-cache gcc g++ make libwebp-dev
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64
WORKDIR /app

# Download Go modules early to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build output.css
RUN chmod +x ./tailwindcss
RUN ./tailwindcss -i public/static/css/input.css -o public/static/css/output.css

# Build templ
RUN chmod +x ./templ
RUN ./templ generate

# Build the application
RUN go build -o server .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache libwebp

# Copy binary from builder stage
COPY --from=builder /app/server .

# Copy static
COPY --from=builder /app/public/ ./public/

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

# Expose port (match your Gin app’s listening port)
EXPOSE 8080


# Run the Go Gin server
ENTRYPOINT ["./entrypoint.sh"]
