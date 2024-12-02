# Use the official Golang image for building the binary
FROM golang:1.23.2 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum, then install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container
COPY .. .

# Build the Go binary for the service, targeting Linux amd64
ARG CMD_PATH
COPY ./${CMD_PATH}/* ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

# Use a minimal image to run the binary
FROM gcr.io/distroless/base-debian11

# Copy the compiled binary from the builder
COPY --from=builder /main /

# Set the entrypoint for the container
ENTRYPOINT ["/main"]