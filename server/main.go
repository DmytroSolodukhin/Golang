package main

import (
	db "port/server/modules/db_manager"
	grpc "port/server/modules/grpc_manager"
	"fmt"
	"flag"
)

const (
	MongoDBHost = "mongodb://"
	PortTOConnect = ":50051"
	DbName = "portsdb"
	defaultDBAddress = "localhost:27017"
	envDBAddress = "db_address"
)

func main() {
	fmt.Println("Starting Server!")
	dbAddress := flag.String(envDBAddress, defaultDBAddress, "")
	mongoHost := MongoDBHost + *dbAddress

	repository := db.ConnectToDB(mongoHost, DbName)
	grpc.StartGRPCServer(PortTOConnect, repository)
}