package main

import (
	db "port/server/modules/db_manager"
	grpc "port/server/modules/grpc_manager"
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
	grpc.StartGRPCServer(PortTOConnect, repository)
}