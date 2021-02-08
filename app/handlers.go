package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	Address string `json:"address" xml:"address"`
	Pin     int    `json:"pin_code" xml:"pin"`
}

func greet(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "Hello Sanvidhan!")
}

func getAllCustomers(w http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{
			"Sanvidhan",
			"Pune",
			411014,
		},
		{
			"Suraj",
			"Pune",
			411014,
		},
		{
			"Aniket",
			"Pune",
			411014,
		},
		{
			"Rahul",
			"Pune",
			411014,
		},
	}

	if request.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}else{
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
