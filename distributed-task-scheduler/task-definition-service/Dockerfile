FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o task-definition-service

CMD ["./task-definition-service"]
