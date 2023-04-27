package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/vbelouso/coffeshop/db"
	"log"
	"net/http"
	"time"
)

const (
	host     string = "localhost"
	port     int    = 5432
	user     string = "postgres"
	password string = "coffeeshop"
	dbname   string = "coffeeshop"
)

type application struct {
	q *db.Queries
}

func main() {
	ctx := context.Background()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	app := &application{q: queries}
	srv := &http.Server{
		Addr:         *addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      app.routes(),
	}
	log.Printf("Starting server on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func openDB(dsn string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	return conn, conn.Ping(context.Background())
}
