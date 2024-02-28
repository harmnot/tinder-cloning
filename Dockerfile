ARG port=9191

# Start from the latest golang base image
FROM golang:1.22.0-alpine

# Add Maintainer Info
LABEL maintainer="Your Name <muhammad@myDomainHere.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o tinder-cloning .

# Expose port 9191 to the outside
EXPOSE $port

# Run the binary program produced by `go install`
CMD ["./tinder-cloning"]