package app

import (
	"Banking/dto"
	"Banking/logger"
	"Banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ah *AccountHandlers) NewAccount(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		writeResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := ah.service.NewAccount(request)
		if appError != nil {
			writeResponse(writer, appError.Code, appError.Message)
		} else {
			writeResponse(writer, http.StatusCreated, account)
		}
	}
}

func (ah *AccountHandlers) MakeTransaction(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	logger.Info("ACCOUNT ID: " + accountId + " | CUSTOMER ID:" + customerId)

	var request dto.TransactionRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		writeResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId

		account, appErr := ah.service.MakeTransaction(request)
		if appErr != nil {
			writeResponse(writer, appErr.Code, appErr.AsMessage())
		} else {
			writeResponse(writer, http.StatusOK, account)
		}
	}
}
