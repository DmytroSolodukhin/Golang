package main

import (
	api "github.com/kazak/Golang/modules/grpcapi"
	restapi "github.com/kazak/Golang/modules/restapi"
	"log"
	"fmt"
	"os"
	"google.golang.org/grpc"
)

const (
	selfHost = ":9090"
	sendAddressDefault = "localhost"
	sendPort = ":50051"
	envSendAddress = "SEND_ADDRESS"
)

func main() {
	fmt.Println("Starting client!")

	sendAddress := os.Getenv(envSendAddress)
	if sendAddress == "" {
		sendAddress = sendAddressDefault
	}
	sendAddress += sendPort

	conn, err := grpc.Dial(sendAddress, grpc.WithInsecure())
	if  err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcClient := api.NewPortServiceClient(conn)

	restapi.Start(selfHost, grpcClient)
}