# syntax=docker/dockerfile:1

# Use the official Golang image from the Docker Hub
# This is based on Alpine Linux and hence is very lightweight
FROM golang:1.21-alpine3.19 as build-stage

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY cmd ./cmd
COPY internal ./internal

# Build the Go app for a Linux environment
# CGO_ENABLED=0 disables C dependencies for the app
# GOOS=linux sets the target operating system to Linux
# The resulting binary is output to /go-app
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app /app/cmd/main.go

FROM scratch
COPY --from=build-stage /go-app /go-app
COPY config.yaml .

EXPOSE 5555

# Run the binary when the container starts
CMD ["/go-app"]
