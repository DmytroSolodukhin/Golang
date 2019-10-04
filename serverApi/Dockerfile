FROM golang:latest
RUN mkdir -p /go/src
WORKDIR /go/src
COPY . /go/src
RUN export GOPATH=/go/src
RUN set -x
RUN chmod +x /go/src/entrypoint.sh
ENTRYPOINT ["/go/src/entrypoint.sh"]