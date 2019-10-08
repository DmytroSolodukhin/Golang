package grpc

import (
	"context"
	api "github.com/kazak/Golang/modules/grpcapi"
	"google.golang.org/grpc"
	"testing"
	"time"
)

type repositoryTest struct {}
func (rt repositoryTest ) GetByID(PortID string) *api.Port {
	return &api.Port{PortId: PortID}
}

func (rt repositoryTest ) GetAll() []*api.Port {
	resp := []*api.Port{
		&api.Port{PortId: "test"},
	}
	return resp

}
func (rt repositoryTest ) Save(*api.Port) bool {
	return true
}
func (rt repositoryTest ) Delete(PortID string) bool {
	return true
}

func TestStartGRPCSsrver(t *testing.T)  {
	mockRepo := &repositoryTest{}
	go StartGRPCServer(":50052", mockRepo)
	time.Sleep(60)
	Convey("gRPC server should be getting request corretly", t, func() {
		conn, _ := grpc.Dial("50052", grpc.WithInsecure())
		defer conn.Close()

		grpcClient := service{repo: mockRepo}

		Convey("Getting singl port", func() {
			req := &api.Request{PortId: "test"}
			exp := &api.Port{PortId: "test"}
			response, _ := grpcClient.Get(context.Background(), req)
			So(response.Port, ShouldResemble, exp)
		})
		Convey("Getting many ports", func() {
			req := &api.Request{PortId: "test"}
			exp := &api.Response{Ports: []*api.Port{
				&api.Port{PortId: "test"},
			}}
			response, _ := grpcClient.GetAll(context.Background(), req)
			So(response, ShouldResemble, exp)
		})
		Convey("Getting save port", func() {
			req := &api.Request{PortId: "test"}
			response, _ := grpcClient.Post(context.Background(), req)
			So(&api.Response{Done: true}, ShouldResemble, response)
		})
		Convey("Getting delete port", func() {
			req := &api.Request{PortId: "test"}
			response, _ := grpcClient.Delete(context.Background(), req)
			So(&api.Response{Done: true}, ShouldResemble, response)
		})
	})
}