FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o database-service

CMD ["./database-service"]
