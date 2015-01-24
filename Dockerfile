FROM golang:1.4.1-wheezy

VOLUME /go/src/app
WORKDIR /go/src/app

CMD go run parser.go
