FROM golang:latest
RUN mkdir -p /go/src/Ports/clientApi
WORKDIR /go/src/Ports/clientApi
COPY . /go/src/Ports/clientApi
RUN export GOPATH=/go/src/Ports/clientApi
RUN set -x
RUN go get -v -u google.golang.org/grpc
RUN go get -v -u github.com/go-chi/chi
RUN go get -u golang.org/x/lint/golint
RUN chmod +x /go/src/Ports/clientApi/entrypoint.sh
ENTRYPOINT ["/go/src/Ports/clientApi/entrypoint.sh"]
