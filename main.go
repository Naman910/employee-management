package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/employees", CreateEmployeeHandler).Methods("POST")
	r.HandleFunc("/employees/{id}", getEmployeeHandler).Methods("GET")
	r.HandleFunc("/employees/{id}", updateEmployeeHandler).Methods("PUT")
	r.HandleFunc("/employees/{id}", deleteEmployeeHandler).Methods("DELETE")
	r.HandleFunc("/employees", listEmployeesHandler).Methods("GET")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
