package main

import (
	api "port/client/models/grpcapi"
	restapi "port/client/models/rest_api"
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