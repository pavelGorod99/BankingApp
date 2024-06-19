package app

import (
	"Banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(writer http.ResponseWriter, request *http.Request) {

	status := request.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomers(status)

	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomerById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomerById(id)
	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customer)
	}
}

func writeResponse(writer http.ResponseWriter, code int, data interface{}) {
	writer.Header().Add("Content-Type", "application/json:")
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		panic(err)
	}
}
