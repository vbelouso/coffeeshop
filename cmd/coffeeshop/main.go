package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vbelouso/coffeshop/internal/config"
	"github.com/vbelouso/coffeshop/internal/handler"
	"github.com/vbelouso/coffeshop/internal/repository"
	"github.com/vbelouso/coffeshop/internal/service"
)

//go:generate swagger generate spec --scan-models -o docs/swagger.yaml
func main() {
	cfg := config.InitializeConfig()

	pg, err := repository.NewPostgreSQLOrderRepository(cfg.DbDSN)
	if err != nil {
		log.Fatal(err)
	}

	svc := service.NewOrderService(pg)
	handler := handler.NewOrderHandler(svc)
	router := mux.NewRouter()
	router.HandleFunc("/orders/{id:[0-9]+}", handler.GetOrderByID).Methods(http.MethodGet)
	srv := &http.Server{
		Addr:         cfg.ServerPort,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
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

// func openDB(dsn string) (*pgxpool.Pool, error) {
// 	ctx := context.Background()
// 	config, err := pgxpool.ParseConfig(dsn)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse database configuration: %w", err)
// 	}

// 	pool, err := pgxpool.NewWithConfig(ctx, config)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create database pool: %w", err)
// 	}

// 	err = pool.Ping(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to establish database connection: %w", err)
// 	}

// 	return pool, nil
// }
