# Use the official golang image from the Docker Hub
FROM golang:1.16 as builder

# Install the Protocol Compiler (protoc)
RUN apt-get update && apt-get install -y protobuf-compiler

# Copy the task.proto file into the Docker image
COPY task.proto .

# Generate the Go code from the task.proto file
RUN protoc --go_out=. task.proto

# Use a new golang image to only have the generated code
FROM golang:1.16 as runtime

# Copy the generated Go code from the previous stage
COPY --from=builder /go/task.pb.go .

# This image can now be used as a base for other images that need the generated code
