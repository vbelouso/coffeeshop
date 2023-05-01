package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/vbelouso/coffeshop/db"
	"log"
	"net/http"
	"time"
)

type application struct {
	q *db.Queries
	c *Config
}

func main() {
	cfg := InitializeConfig()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	conn, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	queries := db.New(conn)

	app := &application{q: queries, c: cfg}
	srv := &http.Server{
		Addr:         cfg.ServerPort,
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
