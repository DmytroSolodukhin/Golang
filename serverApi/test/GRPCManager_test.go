package test

import (
	"Ports/serverApi/model"
	"context"
	grpc "google.golang.org/grpc"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCall(t *testing.T)  {
	Convey("gRPC server should be getting request corretly", t, func() {
		model.StartGRPCSErver(":50052")
		conn, _ := grpc.Dial(":50052", grpc.WithInsecure())
		defer conn.Close()
		grpcClient := model.NewPortServiceClient(conn)

		Convey("Getting chanks", func() {
			request := &model.Request{Method: "GET", PortID: "portID"}
			response, _ := grpcClient.Call(context.Background(), request)
			So(1, ShouldEqual, response)
		})
	})
}