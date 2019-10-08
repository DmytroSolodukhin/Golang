package main

import (
	api "port/client/models/grpcapi"
	restapi "port/client/models/rest_api"
	"log"
	"fmt"
	"flag"
	"google.golang.org/grpc"
)

const (
	selfHost = ":9090"
	sendAddressDefault = "localhost:50051"
	envSendAddress = "SEND_ADDRESS"
)

func main() {
	fmt.Println("Starting client!")

	sendAddress := flag.String(envSendAddress, sendAddressDefault,"server host")

	conn, err := grpc.Dial(*sendAddress, grpc.WithInsecure())
	if  err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcClient := api.NewPortServiceClient(conn)

	restapi.Start(selfHost, grpcClient)
}