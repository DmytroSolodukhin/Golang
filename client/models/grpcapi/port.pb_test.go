package port
//
//import (
//	"Ports/clientApi/model"
//	"context"
//	"google.golang.org/grpc/test/bufconn"
//	grpc "google.golang.org/grpc"
//	"log"
//	"net"
//	"testing"
//	"time"
//	. "github.com/smartystreets/goconvey/convey"
//)
//
//var lis *bufconn.Listener
//
//type PortServiceServerMock struct {
//	model.PortServiceServer
//	seaPort []*model.Port
//}
//
//func (s *PortServiceServerMock) Call(ctx context.Context, request *model.Request) (*model.Response, error) {
//	return &model.Response{Done: true, Port: request.Port}, nil
//}
//
//func init() {
//	lis = bufconn.Listen(1024)
//	server := grpc.NewServer()
//	model.RegisterPortServiceServer(server, &PortServiceServerMock{})
//
//	go func() {
//		if err := server.Serve(lis); err != nil {
//			log.Fatalf("Server exited with error: %v", err)
//		}
//	}()
//}
//
//func bufDialer(string, time.Duration) (net.Conn, error) {
//	conn, err := lis.Dial()
//	return conn, err
//}
//
//func TestSendPort(t *testing.T) {
//	t.Parallel()
//	Convey("Object send and get response corretly", t, func() {
//		ctx := context.Background()
//		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
//		if err != nil {
//			t.Fatalf("Failed to dial bufnet: %v", err)
//		}
//		defer conn.Close()
//		client := model.NewPortServiceClient(conn)
//		objectChank := make(chan *model.Port)
//
//		go func() {
//			objectPort := &model.Port{
//				PortID: "Test",
//				Name: "test",
//				City: "test",
//				Country: "test",
//				Alias: []string{},
//				Regions: []string{},
//				Coordinates: []float64{55.5136433, 25.4052165},
//				Province: "test",
//				Timezone: "test",
//				Unlocs: []string{"TEST"},
//				Code: "52000",
//			}
//
//			objectChank <- objectPort
//		}()
//		Convey("Getting chanks", func() {
//			go func() {
//				time.Sleep(2 * time.Second)
//				close(objectChank)
//			}()
//			resp := model.SendPortToDomainService(client, objectChank)
//			So(1, ShouldEqual, resp)
//		})
//	})
//}
