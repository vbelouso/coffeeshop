package main

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vbelouso/coffeshop/db"
	"log"
	"net/http"
	"time"
)

type application struct {
	q *db.Queries
	c *Config
}

//go:generate swagger generate spec --scan-models -o docs/swagger.yaml
func main() {
	cfg := InitializeConfig()
	InitAPM(cfg.SentryDsn, cfg.Environment)
	defer sentry.Flush(2 * time.Second)

	//logger, err := ConfigureLogging(false)
	//
	//if err != nil {
	//	logger.Fatal("unable to initialize logger", zap.Error(err))
	//}

	conn, err := openDB(cfg.DbDSN)
	if err != nil {
		log.Fatal("unable to establish database connection", err)
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
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database configuration: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to establish database connection: %w", err)
	}

	return pool, nil
}
