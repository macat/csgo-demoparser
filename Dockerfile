FROM golang:1.4.1-wheezy

RUN apt-get update
RUN apt-get install -y bzip2 g++ make wget
RUN wget https://github.com/google/protobuf/releases/download/v2.6.1/protobuf-2.6.1.tar.bz2
RUN tar -xjf protobuf-2.6.1.tar.bz2
RUN cd protobuf-2.6.1 && ./configure --prefix=/usr/local && make && make check && make install && ldconfig

RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go

VOLUME /go/src/app
WORKDIR /go/src/app

CMD go run parser.go
