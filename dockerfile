FROM golang:1.22-alpine  AS build-stage

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY ./ ./

# Build
RUN go build -o /main

# Deploy the application binary into a lean image
FROM alpine:3.19.1 AS build-release-stage

# Update to get latest versions / security
RUN apk update 

# Set destination for COPY
WORKDIR /app

# Copy the application
COPY --from=build-stage /main /app/main

# Run as nonroot
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
USER nonroot:nonroot

ENTRYPOINT ["/app/main"]