FROM golang:latest
RUN mkdir -p /go/src
WORKDIR /go/src
COPY . /go/src
RUN export GOPATH=/go/src
RUN set -x
RUN go build /go/src/main.go
RUN SEND_ADDRESS=port_domain_service /go/src/main