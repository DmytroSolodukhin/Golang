package main

import (
	api "github.com/kazak/Golang/modules/grpcapi"
	restapi "github.com/kazak/Golang/modules/restapi"
	"log"
	"fmt"
	"google.golang.org/grpc"
)

const (
	selfHost = ":9090"
	sendAddress = "localhost:50051"//port_domain_service
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