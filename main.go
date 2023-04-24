package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

const (
	host     string = "localhost"
	port     int    = 5432
	user     string = "postgres"
	password string = "coffeeshop"
	dbname   string = "coffeeshop"
)

var conn *pgx.Conn

func openDB(dsn string) error {
	var err error
	conn, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		return err
	}

	if err = conn.Ping(context.Background()); err != nil {
		return err
	}
	return nil
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	if err := openDB(dsn); err != nil {
		log.Fatal(err)
	}

	// defer conn.Close(context.Background())

	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods(http.MethodGet)
	router.HandleFunc("/orders", getAllOrders).Methods(http.MethodGet)
	router.HandleFunc("/orders/{id:[0-9]+}", getOrder).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", getCustomer).Methods(http.MethodGet)
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
