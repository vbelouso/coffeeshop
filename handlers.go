package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

func getAllOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get all orders")
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	order := struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		ID:   id,
		Name: "fake name",
	}

	err := writeJSON(w, http.StatusOK, envelope{"order": order}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
