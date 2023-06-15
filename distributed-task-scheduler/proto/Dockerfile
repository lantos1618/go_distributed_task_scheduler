# Use the official golang image from the Docker Hub
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /proto

# Install the Protocol Compiler (protoc)
RUN apt-get update && apt-get install -y protobuf-compiler golang-goprotobuf-dev

# Copy the task.proto file into the Docker image
COPY task.proto .

# Generate the Go code from the task.proto file
CMD protoc --go_out=. task.proto

# how to use 
# docker build -t distributed-task-scheduler/proto .
# docker run -it --rm --name proto distributed-task-scheduler/proto 