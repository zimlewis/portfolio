# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# Build frontend asset
FROM node:latest AS asset-build-stage
WORKDIR /app
COPY package*.json .

RUN npm install
COPY . .
RUN npx vite build

# Build server
FROM golang:1.25-alpine AS build-stage

RUN apk add --no-cache gcc g++ make libwebp-dev
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum .
RUN go mod download

COPY --from=generate-stage /app .
RUN go build -o server .

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache libwebp

# Copy binary from builder stage
COPY --from=build-stage /app/server .
COPY --from=asset-build-stage /app/public ./public

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]


