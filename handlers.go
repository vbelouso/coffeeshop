package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "not found", err)
		return
	}

	if id < 1 {
		fmt.Fprintf(os.Stderr, "not found", err)
		return
	}

	query := `
	SELECT CUSTOMER_ID,
	FULL_NAME,
	CREATED_AT
	FROM PUBLIC.CUSTOMERS
	WHERE CUSTOMER_ID = $1;`

	customer := &Customer{}

	err = conn.QueryRow(context.Background(), query, id).Scan(
		&customer.ID,
		&customer.FullName,
		&customer.Created)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	err = writeJSON(w, http.StatusOK, envelope{"customer": customer}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
