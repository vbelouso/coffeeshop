package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods(http.MethodGet)
	router.HandleFunc("/orders", getAllOrders).Methods(http.MethodGet)
	router.HandleFunc("/orders/{id:[0-9]+}", getOrder).Methods(http.MethodGet)
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
