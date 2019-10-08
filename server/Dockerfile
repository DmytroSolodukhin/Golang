FROM golang:latest
RUN mkdir -p /go/src
WORKDIR /go/src
COPY . /go/src
RUN export GOPATH=/go/src
RUN set -x
RUN go get -u github.com/kazak/Golang/modules/db_manager
RUN go get -u github.com/kazak/Golang/modules/grpc_manager
RUN go get -u github.com/kazak/Golang/modules/grpcapi
RUN go build /go/src/main.go
ENTRYPOINT ["DB_ADDRESS=mongo:27017", "/go/src/main"]
