package main

import (
	api "github.com/kazak/grpcapi"
	restapi "github.com/kazak/restapi"
	"log"
	"fmt"
	"google.golang.org/grpc"
)

const (
	selfHost = ":9090"
	sendAddress = "port_domain_service:50051"
)

func main() {
	fmt.Println("Starting client!")
	conn, err := grpc.Dial(sendAddress, grpc.WithInsecure())
	if  err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcClient := api.NewPortServiceClient(conn)

	restapi.Start(selfHost, grpcClient)
}