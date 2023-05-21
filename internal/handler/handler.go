package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vbelouso/coffeshop/internal/helpers"
	"github.com/vbelouso/coffeshop/internal/service"
)

type orderHandler struct {
	service service.IOrderService
}

func NewOrderHandler(service service.IOrderService) *orderHandler {
	return &orderHandler{service: service}
}

func (h *orderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}

	order, err := h.service.GetOrderByID(int(id))
	if err != nil {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"customer": order}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
