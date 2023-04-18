package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods(http.MethodGet)
	router.HandleFunc("/orders", getAllOrders).Methods(http.MethodGet)
	router.HandleFunc("/orders/{id}", getOrder).Methods(http.MethodGet)
	http.Handle("/", router)

	srv := &http.Server{
		Addr:         *addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	log.Printf("Starting server on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe(), router)
}

func homePage(w http.ResponseWriter, r *http.Request) {
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

	payload, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
