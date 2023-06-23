FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o cli-user-interface

CMD ["./cli-user-interface"]
