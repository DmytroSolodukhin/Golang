package main

import (
	db "github.com/kazak/repository"
	grpc "github.com/kazak/grpcmanager"
	"fmt"
)

const (
	MongoDBHost = "mongodb://mongo:27017"
	PortTOConnect = ":50051"
	DbName = "portsdb"
)

func main() {
	fmt.Println("Starting Server!")
	repository := db.ConnectToDB(MongoDBHost, DbName)
	grpc.StartGRPCSErver(PortTOConnect, repository)
}