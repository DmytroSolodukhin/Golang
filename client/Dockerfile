FROM golang:latest
RUN mkdir -p /go/src
WORKDIR /go/src
COPY . /go/src
RUN export GOPATH=/go/src
RUN set -x
RUN go build /go/src/main.go
RUN go get -u github.com/kazak/Golang/modules/grpcapi
RUN go get -u github.com/kazak/Golang/modules/port_factory
RUN go get -u github.com/kazak/Golang/modules/restapi
ENTRYPOINT ["SEND_ADDRESS=port_domain_service", "/go/src/main"]
