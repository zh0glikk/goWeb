FROM golang:1.14

WORKDIR /go/src/goWeb
COPY . .

RUN go build -o main .

CMD ["./main"]