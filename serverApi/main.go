package main

import (
	db "github.com/kazak/Golang/modules/repository"
	grpc "github.com/kazak/Golang/modules/grpcmanager"
	"fmt"
	"os"
)

const (
	MongoDBHost = "mongodb://"
	PortTOConnect = ":50051"
	DbName = "portsdb"
	defaultDBAddress = "mongo:27017"
	envDBAddress = "DB_ADDRESS"
)

func main() {
	fmt.Println("Starting Server!")
	dbAddress := os.Getenv(envDBAddress)
	if dbAddress == "" {
		dbAddress = defaultDBAddress
	}
	mongoHost := MongoDBHost + dbAddress

	repository := db.ConnectToDB(mongoHost, DbName)
	grpc.StartGRPCServer(PortTOConnect, repository)
}