FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o task-scheduler-service

CMD ["./task-scheduler-service"]
