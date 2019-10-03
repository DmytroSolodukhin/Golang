FROM golang:latest
RUN mkdir -p /go/src/Ports/serverApi
WORKDIR /go/src/Ports/serverApi
COPY . /go/src/Ports/serverApi
RUN export GOPATH=/go/src/Ports/serverApi
RUN set -x
RUN go get -v -u google.golang.org/grpc
RUN go get -v -u go.mongodb.org/mongo-driver/mongo
RUN go get -u golang.org/x/lint/golint
RUN chmod +x /go/src/Ports/serverApi/entrypoint.sh
ENTRYPOINT ["/go/src/Ports/serverApi/entrypoint.sh"]