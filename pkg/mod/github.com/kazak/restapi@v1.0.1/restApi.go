package restapi

import (
	api "github.com/kazak/grpcapi"
	"github.com/kazak/data"
	"bufio"
	"context"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi"
)

type response struct {
	data      []byte
	message   string
	error     bool
	code      int
}

var grpcClient api.PortServiceClient

func (response *response) output(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	out, _ := json.Marshal(response)
	res.Write(out)
}

// Getting all ports. (static limit 100)
func gatAll(res http.ResponseWriter, req *http.Request) {
	response := &response{
		code: 200,
		error: false,
		message: "",
	}
	ports, err := grpcClient.GetAll(context.Background(), &api.Request{})
	if err != nil {
		response.code = 502
		response.error = true
		response.message = "Server lost connect"
	}
	jsonData, err := json.Marshal(ports.Ports)
	if err != nil {
		response.code = 500
		response.error = true
		response.message = "Can't scan data"
	}
	res.Write(jsonData)
}

// Getting port by id
func getByID(res http.ResponseWriter, req *http.Request)  {
	response := &response{
		code: 200,
		error: false,
		message: "",
	}
	portID := chi.URLParam(req, "portID")
	request := &api.Request{PortId: portID}
	out, err := grpcClient.Get(context.Background(), request)
	if err != nil {
		response.code = 502
		response.error = true
		response.message = "Server lost connect"
	}
	response.data, _ = json.Marshal(out.Port)
	if err != nil {
		response.code = 402
		response.error = true
		response.message = "Can't serislize"
	}
	response.output(res)
}

// Create and save ports
func post(res http.ResponseWriter, req *http.Request) {
	response := &response{
		code: 200,
		error: false,
		message: "",
	}
	chankText, chankPort := make(chan string), make(chan *api.Port)
	scanner := bufio.NewScanner(req.Body)
	defer req.Body.Close()

	if err := scanner.Err(); err != nil {
		response.code = 500
		response.error = true
		response.message = "Can't scan data"
	}
	go data.ScanRequestBodyToChank(scanner, chankText)
	go data.StartProdactionPort(chankText, chankPort)

	countOfSavedPort := sendPortToDomainService(grpcClient, chankPort)
	response.message = string(countOfSavedPort)
	response.output(res)
}

// Remove ports
func remove(res http.ResponseWriter, req *http.Request) {
	response := &response{
		code: 200,
		error: false,
		message: "",
	}
	portID := chi.URLParam(req, "portID")
	request := &api.Request{PortId: portID}
	output, _ := grpcClient.Delete(context.Background(), request)

	if !output.Done {
		response.code = 402
		response.error = true
		response.message = "Can't dalete data"
	}
	response.output(res)
}

// Start REST Api
func Start(selfHost string, client api.PortServiceClient)  {
	grpcClient = client
	rout := chi.NewRouter()
	rout.Get("/", gatAll)
	rout.Get("/{portID}", getByID)
	rout.Post("/", post)
	rout.Delete("/{portID}", remove)
	http.ListenAndServe(selfHost, rout)
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

