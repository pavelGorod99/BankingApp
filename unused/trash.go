package unused

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

//router.HandleFunc("/greet", greet).Methods(http.MethodGet)
//router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
//router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

func createCustomer(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "POST REQUEST RECEIVED")
}

func getCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprint(writer, vars["customer_id"])
}

// FIND ALL FROM DATABASE
//for rows.Next() {
//	var c Customer
//	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
//
//	if err != nil {
//		logger.Error("Error while scanning customers " + err.Error())
//		return nil, errs.NewUnexpectedError("Unexpected database error")
//	}
//	customers = append(customers, c)
//}
