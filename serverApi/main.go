package main

import (
	db "github.com/kazak/Golang/modules/db_manager"
	grpc "github.com/kazak/Golang/modules/grpc_manager"
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