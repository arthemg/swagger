FROM golang
MAINTAINER Artsem Holdvekht

ENV GOBIN $GOPATH/bin

ADD . /go/src/github.com/arthemg/dataParser
RUN go install /go/src/github.com/arthemg/dataParser/cmd/data-parser-server
WORKDIR /go/src/github.com/arthemg/dataParser
ENTRYPOINT /go/bin/data-parser-server --port 8090 --host 0.0.0.0

# serving HTTP of 8090
EXPOSE 8090