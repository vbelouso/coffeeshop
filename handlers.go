package main

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	err = writeJSON(w, http.StatusOK, envelope{"customer": c}, nil)
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
	err = writeJSON(w, http.StatusOK, envelope{"customers": c}, nil)
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

	err = writeJSON(w, http.StatusOK, envelope{"order": o}, nil)
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
	err = writeJSON(w, http.StatusOK, envelope{"orders": c}, nil)
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

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		//The base64-encoded public key downloaded from Keycloak.
		//TODO "REPLACE_WITH_ENV"
		base64EncodedPublicKey := "REPLACE_WITH_ENV"
		publicKey, err := parseRSAPublicKey(base64EncodedPublicKey)
		if err != nil {
			panic(err)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return publicKey, nil
		})
		if err != nil {
			fmt.Println("Error parsing or validating token:", err)
			return
		}

		if !token.Valid {
			fmt.Println("Invalid token")
			return
		}

		//claims := token.Claims.(jwt.MapClaims)
		//fmt.Println("Claims:", claims)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
