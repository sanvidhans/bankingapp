package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sanvidhans/bankingapp/dto"
	"github.com/sanvidhans/bankingapp/logger"
	"github.com/sanvidhans/bankingapp/service"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	Address string `json:"address" xml:"address"`
	Pin     int    `json:"pin_code" xml:"pin"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

type AccountHandlers struct {
	service service.AccountService
}

func (ah *AccountHandlers) createAccount(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}else{
		request.CustomerId = customerId
		h, err := ah.service.NewAccount(request)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		}else{
			writeResponse(w, http.StatusCreated, h)
		}}

}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}else{
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)

	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}else {
		writeResponse(w, http.StatusOK, customer)
	}

}

func writeResponse(w http.ResponseWriter, code int, data interface{}){
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Error("Could not server the response!")
		panic(err)
	}
}