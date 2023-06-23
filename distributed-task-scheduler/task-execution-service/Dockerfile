FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o task-execution-service

CMD ["./task-execution-service"]
