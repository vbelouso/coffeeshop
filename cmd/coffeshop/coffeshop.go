package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vbelouso/coffeshop/internal/config"
	"github.com/vbelouso/coffeshop/internal/db"
	"log"
	"net/http"
	"time"
)

type application struct {
	q *db.Queries
	c *config.Config
}

//go:generate swagger generate spec --scan-models -o docs/swagger.yaml
func main() {
	cfg := config.InitializeConfig()
	conn, err := openDB(cfg.DBDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

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

func openDB(dsn string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	defer pool.Close()
	return pool, pool.Ping(ctx)
}
