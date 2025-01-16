FROM golang:1.23-alpine3.21
# Set the Current Working Directory inside the container
WORKDIR /
# Copy go.mod and go.sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Set environment variable for the views directory
RUN go build
# Expose port 8080 to the outside world
EXPOSE 8888
# Set environment variable for Gin mode
ENV GIN_MODE=release
# Run the executable
CMD ["./globalbans"]
