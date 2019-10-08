package restapi

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/kazak/data"
	api "github.com/kazak/grpcapi"
	"net/http"
)

type response struct {
	data      string
	code      int
}

var grpcClient api.PortServiceClient

func (response *response) output(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(response.code)

	res.Write([]byte(response.data))
}

// Getting all ports. (static limit 100)
func getAll(res http.ResponseWriter, req *http.Request) {
	response := &response{
		code: 200,
	}
	out, err := grpcClient.GetAll(context.Background(), &api.Request{})

	if err != nil {
		response.code = 502
	}
	jsonData, _ := json.Marshal(out.Ports)
	response.data = string(jsonData)

	if err != nil {
		response.code = 500
	}

	response.output(res)
}

// Getting port by id
func getByID(res http.ResponseWriter, req *http.Request)  {
	response := &response{
		code: 200,
	}
	portID := chi.URLParam(req, "portID")
	request := &api.Request{PortId: portID}
	out, err := grpcClient.Get(context.Background(), request)

	if err != nil {
		response.code = 502
	}
	jsonData, _ := json.Marshal(out.Port)
	response.data = string(jsonData)

	if err != nil {
		response.code = 402
	}

	response.output(res)
}

// Create and save ports
func post(res http.ResponseWriter, req *http.Request) {
	response := &response{
		code: 200,
	}
	chankText, chankPort := make(chan string), make(chan *api.Port)
	scanner := bufio.NewScanner(req.Body)
	defer req.Body.Close()

	if err := scanner.Err(); err != nil {
		response.code = 500
	}

	go data.ScanRequestBodyToChank(scanner, chankText)
	go data.StartProdactionPort(chankText, chankPort)

	countOfSavedPort := sendPortToDomainService(grpcClient, chankPort)
	response.data = string(countOfSavedPort)

	response.output(res)
}

// Remove ports
func remove(res http.ResponseWriter, req *http.Request) {
	response := &response{
		code: 200,
	}
	portID := chi.URLParam(req, "portID")
	request := &api.Request{PortId: portID}
	output, _ := grpcClient.Delete(context.Background(), request)

	if !output.Done {
		response.code = 402
	} else {
		response.data = "true"
	}

	response.output(res)
}

// Start REST Api
func Start(selfHost string, client api.PortServiceClient) {
	fmt.Println("Ssfsdfsdfsdf!")

	grpcClient = client
	rout := chi.NewRouter()
	rout.Post("/", post)
	rout.Get("/", getAll)
	rout.Get("/{portID}/", getByID)
	rout.Delete("/{portID}/", remove)
	fmt.Println("@#$@#$!")

	_ = http.ListenAndServe(selfHost, rout)
}

// SendPortToDomainService send port object to service.
func sendPortToDomainService(grpcClient api.PortServiceClient, chPort <-chan *api.Port) int {
	var quantity = 0
	for {
		portObject, done := <- chPort

		if !done {
			return quantity
		}
		request := &api.Request{Port: portObject}
		response, _ := grpcClient.Post(context.Background(), request)

		if response.Done {
			quantity++
		}
	}
}

