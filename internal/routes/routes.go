package routes

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func (a *application) routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", a.home).Methods(http.MethodGet)
	router.HandleFunc("/orders", a.getAllOrders).Methods(http.MethodGet)
	router.HandleFunc("/orders/{id:[0-9]+}", a.getOrder).Methods(http.MethodGet)
	router.HandleFunc("/customers", a.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", a.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/healthz", a.healthzHandler).Methods(http.MethodGet)
	opts := middleware.SwaggerUIOpts{SpecURL: "/docs/swagger.yaml"}
	router.Handle("/docs", middleware.SwaggerUI(opts, nil))
	//router.HandleFunc("/docs/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "docs/swagger.yaml")
	//})

	return router
}
