FROM golang:1.8-alpine

MAINTAINER Artsem Holdvekht

ENV GOBIN $GOPATH/bin

RUN go get -u github.com/golang/dep/cmd/dep
ADD . /go/src/github.com/arthemg/dataParser
RUN dep init && dep ensure
RUN go install github.com/arthemg/dataParser/cmd/data-parser-server
# RUN go install /go/src/github.com/arthemg/dataParser/cmd/data-parser-server
WORKDIR /go/src/github.com/arthemg/dataParser
ENTRYPOINT /go/bin/data-parser-server --port=50051 --host 0.0.0.0

# serving HTTP of 8090
EXPOSE 50051