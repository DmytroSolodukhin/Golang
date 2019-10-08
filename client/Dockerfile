FROM golang:latest
RUN mkdir -p /go/src
WORKDIR /go/src
COPY . /go/src
RUN export GOPATH=/go/src
RUN set -x
RUN go build /go/src/main.go
CMD ["/go/src/main", "-send_address=port_domain_service:50051"]
