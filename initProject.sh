#!/bin/bash

# Create directories
mkdir -p distributed-task-scheduler/{task-definition-service,task-scheduler-service,task-execution-service,database-service,cli-user-interface}/{handlers}
mkdir -p distributed-task-scheduler/proto

# Create main.go and Dockerfile in each service directory
for service in task-definition-service task-scheduler-service task-execution-service database-service user-interface
do
    touch distributed-task-scheduler/${service}/main.go
    touch distributed-task-scheduler/${service}/Dockerfile
done

# Create task.proto in proto directory
touch distributed-task-scheduler/proto/task.proto

# Create docker-compose.yml in the root directory
touch distributed-task-scheduler/docker-compose.yml

# Initialize Go modules
cd distributed-task-scheduler
for service in task-definition-service task-scheduler-service task-execution-service database-service cli-user-interface
do
    cd ${service}
    go mod init zug.dev/${service}
    cd ..
done