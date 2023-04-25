package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

func getAllOrders(w http.ResponseWriter, r *http.Request) {
	o, err := getOrders()
	if err != nil {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}
	err = writeJSON(w, http.StatusOK, envelope{"orders": o}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	errorMsg := "resource bla-bla"
	httpStatus := http.StatusNotFound
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err == nil {
		if id >= 0 {
			o, err := getOrderByID(id)
			if err == nil {
				err = writeJSON(w, http.StatusOK, envelope{"order": o}, nil)
				if err == nil {
					return
				}
			}
		}
	}
	log.Println(err)
	http.Error(w, errorMsg, httpStatus)
}

// func getOrder(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(mux.Vars(r)["id"])
// 	if err != nil {
// 		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
// 		return
// 	}
// 	if id < 1 {
// 		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
// 		return
// 	}

// 	o, err := getOrderByID(id)
// 	if err != nil {
// 		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
// 		return
// 	}

// 	err = writeJSON(w, http.StatusOK, envelope{"order": o}, nil)
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
// 	}
// }

func getCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}
	if id < 1 {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}
	c, err := getCustomerByID(id)
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

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	c, err := getCustomers()
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
