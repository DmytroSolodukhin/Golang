package restapi

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	portFactory "github.com/kazak/Golang/modules/port_factory"
	api "github.com/kazak/Golang/modules/grpcapi"
	"net/http"
)

type response struct {
	data      string
	code      int
}

type restRoute interface {
	getAll(res http.ResponseWriter, req *http.Request)
	getByID(res http.ResponseWriter, req *http.Request)
	post(res http.ResponseWriter, req *http.Request)
	remove(res http.ResponseWriter, req *http.Request)
}

type restRout struct {
	grpcClient api.PortServiceClient
}

func (response *response) output(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(response.code)

	res.Write([]byte(response.data))
}

// Getting all ports. (static limit 100)
func (route *restRout) getAll(res http.ResponseWriter, req *http.Request) {
	response := &response{
		code: 200,
	}
	out, err := route.grpcClient.GetAll(context.Background(), &api.Request{})

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
func (route *restRout) getByID(res http.ResponseWriter, req *http.Request)  {
	response := &response{
		code: 200,
	}
	portID := chi.URLParam(req, "portID")
	request := &api.Request{PortId: portID}
	out, err := route.grpcClient.Get(context.Background(), request)

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
func (route *restRout) post(res http.ResponseWriter, req *http.Request) {
	response := &response{
		code: 200,
	}
	chankText, chankPort := make(chan string), make(chan *api.Port)
	scanner := bufio.NewScanner(req.Body)
	defer req.Body.Close()

	if err := scanner.Err(); err != nil {
		response.code = 500
	}

	go portFactory.ScanRequestBodyToChank(scanner, chankText)
	go portFactory.StartProdactionPort(chankText, chankPort)

	countOfSavedPort := route.sendPortToDomainService(chankPort)
	response.data = string(countOfSavedPort)

	response.output(res)
}

// Remove ports
func (route *restRout) remove(res http.ResponseWriter, req *http.Request) {
	response := &response{
		code: 200,
	}
	portID := chi.URLParam(req, "portID")
	request := &api.Request{PortId: portID}
	output, _ := route.grpcClient.Delete(context.Background(), request)

	if !output.Done {
		response.code = 402
	} else {
		response.data = "true"
	}

	response.output(res)
}

// SendPortToDomainService send port object to service.
func (route *restRout) sendPortToDomainService(chPort <-chan *api.Port) int {
	var quantity = 0
	for {
		portObject, done := <- chPort

		if !done {
			return quantity
		}
		request := &api.Request{Port: portObject}
		response, _ := route.grpcClient.Post(context.Background(), request)

		if response.Done {
			quantity++
		}
	}
}

// Start REST Api
func Start(selfHost string, client api.PortServiceClient) {
	restRout := &restRout{
		grpcClient: client,
	}
	rout := chi.NewRouter()
	rout.Post("/", restRout.post)
	rout.Get("/", restRout.getAll)
	rout.Get("/{portID}/", restRout.getByID)
	rout.Delete("/{portID}/", restRout.remove)

	_ = http.ListenAndServe(selfHost, rout)
}

