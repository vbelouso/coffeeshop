package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (a *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

func (a *application) getCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}

	c, err := a.q.GetCustomer(context.Background(), id)
	if err != nil {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"customer": c}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (a *application) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	c, err := a.q.ListCustomers(context.Background())
	if err != nil {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}
	err = writeJSON(w, http.StatusOK, envelope{"customers": c}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (a *application) getOrder(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}

	o, err := a.q.GetOrder(context.Background(), id)
	if err != nil {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"order": o}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (a *application) getAllOrders(w http.ResponseWriter, r *http.Request) {
	c, err := a.q.ListOrders(context.Background())
	if err != nil {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}
	err = writeJSON(w, http.StatusOK, envelope{"orders": c}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (a *application) healthzHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available"}`
	js = fmt.Sprintf(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}
