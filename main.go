package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vbelouso/coffeshop/db"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	q *db.Queries
	c *Config
}

//go:generate swagger generate spec --scan-models -o docs/swagger.yaml
func main() {
	cfg := InitializeConfig()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	conn, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//conn := openDB(dsn)
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

//func openDB(dsn string) (*pgx.Conn, error) {
//	conn, err := pgx.Connect(context.Background(), dsn)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return conn, conn.Ping(context.Background())
//}

func openDB(dsn string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()
	return pool, pool.Ping(ctx)
}
