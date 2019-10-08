package restapi

import (
	"bytes"
	"github.com/go-chi/chi"
	"strconv"

	"encoding/json"
	"io"
	"os"

	api "github.com/kazak/grpcapi"
	"google.golang.org/grpc"
	"net/http"
	"net/http/httptest"
	"testing"
	"context"
	. "github.com/smartystreets/goconvey/convey"
	)

type mockClient struct {}
func (s mockClient) Get(ctx context.Context, in *api.Request, opts ...grpc.CallOption) (*api.Response, error) {
	return &api.Response{Port: &api.Port{PortId: "test"}}, nil
}
func (s mockClient) GetAll(ctx context.Context, in *api.Request, opts ...grpc.CallOption) (*api.Response, error) {
	return &api.Response{Ports: []*api.Port{&api.Port{PortId: "test"}}}, nil
}
func (s mockClient) Post(ctx context.Context, in *api.Request, opts ...grpc.CallOption) (*api.Response, error) {
	return &api.Response{Done: true}, nil
}
func (s mockClient) Delete(ctx context.Context, in *api.Request, opts ...grpc.CallOption) (*api.Response, error) {
	return &api.Response{Done: true}, nil
}

func TestHttpApi(t *testing.T) {
	Convey("REST API should be response corretly", t, func() {
		grpcClient = &mockClient{}

		Convey("Post method", func() {
			file, _ := os.Open("./fixture/test.json")
			defer file.Close()
			body := &bytes.Buffer{}

			io.Copy(body, file)

			req, _ := http.NewRequest("POST", "/", body)
			req.Header.Add("Content-Type", "multipart/form-data")
			req.Header.Add("Content-Length", strconv.Itoa(body.Len()))

			response := httptest.NewRecorder()
			post(response, req)
			response.Body.String()
		})

		Convey("Get method  without id", func() {
			req, _ := http.NewRequest("GET", "/", nil)
			response := httptest.NewRecorder()

			getAll(response, req)
			got := response.Body.String()
			ports := []*api.Port{&api.Port{PortId: "test"}}
			exp, _ := json.Marshal(ports)
			So(got, ShouldResemble, string(exp))
		})

		Convey("Get method  with id", func() {
			req, _ := http.NewRequest("GET", "/test", nil)
			rctx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			response := httptest.NewRecorder()

			getByID(response, req)
			got := response.Body.String()
			So(got, ShouldResemble, "{\"port_id\":\"test\"}")
		})

		Convey("Delete method", func() {
			req, _ := http.NewRequest("DELETE", "/test", nil)
			rctx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			response := httptest.NewRecorder()

			remove(response, req)
			got := response.Body.String()
			So(got, ShouldEqual, "true")
		})
	})
}
