package main

import (
	db "github.com/kazak/repository"
	grpc "github.com/kazak/grpcmanager"
	"fmt"
)

const (
	MongoDBHost = "mongodb://localhost:27017"//mongo
	PortTOConnect = ":50051"
	DbName = "portsdb"
)

func main() {
	fmt.Println("Starting Server!")
	repository := db.ConnectToDB(MongoDBHost, DbName)
	grpc.StartGRPCServer(PortTOConnect, repository)
}