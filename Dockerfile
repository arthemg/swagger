FROM golang:1.11.1-alpine3.8 as builder

LABEL maintainer="Artsem Holdvekht"

WORKDIR /go/src/swagger

COPY . .

RUN apk --no-cache add ca-certificates shared-mime-info mailcap git build-base curl && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN make

# Second stage
FROM alpine:3.8

RUN addgroup -S swagger && \
    adduser -S -G swagger swagger

COPY --from=builder /go/src/swagger/swagger /home/swagger/

RUN chown -R swagger:swagger /home/swagger

USER swagger

ENTRYPOINT ["/home/swagger/swagger", "--port=50051", "--host", "0.0.0.0"]

# serving HTTP of 8090
EXPOSE 50051