FROM golang:alpine

MAINTAINER Artsem Holdvekht

ENV GOBIN $GOPATH/bin
 

RUN apk --no-cache add ca-certificates shared-mime-info mailcap git build-base && \
#   go get -u github.com/asaskevich/govalidator &&\
#   go get -u golang.org/x/net/context &&\
#   go get -u github.com/jessevdk/go-flags &&\
#   go get -u golang.org/x/net/context/ctxhttp &&\
#   go get -u github.com/tatsushid/go-fastping &&\
#   go get -u github.com/go-openapi/runtime &&\
#   go get -u github.com/docker/go-units &&\
#   go get -u github.com/go-openapi/analysis &&\
#   go get -u github.com/go-openapi/loads &&\
#   go get -u github.com/go-openapi/spec &&\
#   go get -u github.com/go-openapi/validate
    go get -u github.com/golang/dep/cmd/dep &&\
    dep init
# RUN dep ensure

ADD . /go/src/github.com/arthemg/dataParser
# RUN dep init && dep ensure
RUN go install github.com/arthemg/dataParser/cmd/data-parser-server &&\ 
dep init && dep ensure
# RUN go install /go/src/github.com/arthemg/dataParser/cmd/data-parser-server
WORKDIR /go/src/github.com/arthemg/dataParser
ENTRYPOINT /go/bin/data-parser-server --port=50051 --host 0.0.0.0

# serving HTTP of 8090
EXPOSE 50051