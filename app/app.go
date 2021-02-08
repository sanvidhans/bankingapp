package app

import (
	"fmt"
	"log"
	"net/http"
)

func Start(){
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	fmt.Println("Server Running on http://hocalhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}