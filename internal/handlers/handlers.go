package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vbelouso/coffeshop/internal/helpers"
	"log"
	"net/http"
	"strconv"
)

func (a *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

// getCustomer return the customer from the coffeeshop store.
func (a *application) getCustomer(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /customers/{id} getCustomer
	//
	// Retrieves the details about customer.
	//
	// Could be any customer
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: Customer ID to filter by
	//   required: true
	//   type: integer
	//   format: int64
	// responses:
	//   '200':
	//     description: Success
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/Customer"
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}

	c, err := a.q.GetCustomer(context.Background(), id)
	if err != nil {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"customer": c}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (a *application) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	c, err := a.q.ListCustomers(context.Background())
	if err != nil {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}
	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"customers": c}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (a *application) getOrder(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}

	o, err := a.q.GetOrder(context.Background(), id)
	if err != nil {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"order": o}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (a *application) getAllOrders(w http.ResponseWriter, r *http.Request) {
	c, err := a.q.ListOrders(context.Background())
	if err != nil {
		http.Error(w, "The requested resource was not found.", http.StatusNotFound)
		return
	}
	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"orders": c}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (a *application) healthzHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available"}`
	js = fmt.Sprintf(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}
