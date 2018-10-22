# FROM golang:alpine

# MAINTAINER Artsem Holdvekht

# ENV GOBIN $GOPATH/bin


# RUN apk --no-cache add ca-certificates shared-mime-info mailcap git build-base && \
#     go get -u github.com/asaskevich/govalidator &&\
#     go get -u golang.org/x/net/context &&\
#     go get -u github.com/jessevdk/go-flags &&\
#     go get -u golang.org/x/net/context/ctxhttp &&\
#     go get -u github.com/tatsushid/go-fastping &&\
#     go get -u github.com/go-openapi/runtime &&\
#     go get -u github.com/docker/go-units &&\
#     go get -u github.com/go-openapi/analysis &&\
#     go get -u github.com/go-openapi/loads &&\
#     go get -u github.com/go-openapi/spec &&\
#     go get -u github.com/go-openapi/validate &&\
#     go get -u github.com/golang/dep/cmd/dep
# # RUN dep ensure

# ADD . /go/src/github.com/arthemg/dataParser
# RUN go install github.com/arthemg/dataParser/cmd/data-parser-server
# # RUN go install /go/src/github.com/arthemg/dataParser/cmd/data-parser-server
# WORKDIR /go/src/github.com/arthemg/dataParser
# ENTRYPOINT /go/bin/data-parser-server --port=50051 --host 0.0.0.0

# # serving HTTP of 8090
# EXPOSE 50051


FROM vmj0/golang-dep:1.11.1-stretch-0.5.0 as build

# RUN apt-get update && apt-get install -y unzip --no-install-recommends && \
#     apt-get autoremove -y && apt-get clean -y && \
#     wget -O dep.zip https://github.com/golang/dep/releases/download/v0.3.0/dep-linux-amd64.zip && \
#     echo '96c191251164b1404332793fb7d1e5d8de2641706b128bf8d65772363758f364  dep.zip' | sha256sum -c - && \
#     unzip -d /usr/bin dep.zip && rm dep.zip

RUN mkdir -p /go/src/github.com/arthemg/dataParser
WORKDIR /go/src/github.com/arthemg/dataParser

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go install -ldflags "-s -w" -a -installsuffix cgo /go/src/github.com/arthemg/dataParser/cmd/data-parser-server

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /go/src/github.com/arthemg/dataParser  .
ENTRYPOINT /go/bin/data-parser-server --port=50051 --host 0.0.0.0

EXPOSE 50051
