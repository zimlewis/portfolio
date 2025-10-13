# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# Generate output.css
FROM alpine:latest AS generate-static
RUN apk add --no-cache gcc g++ make libwebp-dev
WORKDIR /app
COPY . .

RUN wget https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64-musl 
RUN chmod +x ./tailwindcss-linux-x64-musl 
RUN ./tailwindcss-linux-x64-musl -i public/static/css/input.css -o public/static/css/output.css


# Build
FROM golang:1.25-alpine AS build-stage


RUN apk add --no-cache gcc g++ make libwebp-dev
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum .
RUN go mod download

COPY --from=generate-stage /app .


RUN go build -o server .

# Run 
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache libwebp

# Copy binary from builder stage
COPY --from=build-stage /app/server .

# Copy static
COPY --from=generate-static /app/public/ ./public/

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

# Expose port (match your Gin app’s listening port)
EXPOSE 8080


# Run the Go Gin server
ENTRYPOINT ["./entrypoint.sh"]
