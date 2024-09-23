FROM golang:1.23-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main cmd/main.go
RUN chmod +x /app

CMD ["/app/main"]