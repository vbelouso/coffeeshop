package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vbelouso/coffeshop/internal/config"
	"github.com/vbelouso/coffeshop/internal/controller"
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
	handler := controller.NewOrderHandler(svc)
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
