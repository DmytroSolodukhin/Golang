package grpc

import (
	api "github.com/kazak/grpcapi"
	db "github.com/kazak/repository"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type service struct {
	repo *db.Repository
}

func (s *service) Get(ctx context.Context, req *api.Request) (*api.Response, error) {
	PortId := req.PortId
	port := s.repo.GetByID(PortId)
	return &api.Response{Port: port}, nil
}

func (s *service) GetAll(ctx context.Context, req *api.Request) (*api.Response, error) {
	ports := s.repo.GetAll()
	return &api.Response{Ports: ports}, nil
}
func (s *service) Post(ctx context.Context, req *api.Request) (*api.Response, error) {
	port := req.Port
	res := s.repo.Save(port)
	return &api.Response{Done: res}, nil
}
func (s *service) Dalete(ctx context.Context, req *api.Request) (*api.Response, error) {
	PortId := req.PortId
	res := s.repo.Delete(PortId)
	return &api.Response{Done: res}, nil
}

//Start gRPC server
func StartGRPCSErver(PortTOConnect string, repository *db.Repository)  {
	// Start gRPC server and listen tcp
	lis, err := net.Listen("tcp", PortTOConnect)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	api.RegisterPortServiceServer(server, &api.PortServiceServer {repository})

	// Register to response gRPC.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}