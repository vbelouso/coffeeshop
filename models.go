package main

import (
	"errors"
	"time"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Customer struct {
	ID       int       `json:"id,omitempty"`
	FullName string    `json:"fullName,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

// type CustomerModel struct {
// 	conn *pgx.Conn
// }

// func (m CustomerModel) Get(id int) (*Customer, error) {
// 	if id < 1 {
// 		return nil, ErrRecordNotFound
// 	}

// 	query := `
// 		SELECT CUSTOMER_ID,
// 		FULL_NAME,
// 		CREATED_AT
// 		FROM public.customers
// 		WHERE CUSTOMER_ID= ?`

// 	customer := &Customer{}

// 	err := m.conn.QueryRow(context.Background(), query, id).Scan(
// 		&customer.ID,
// 		&customer.FullName,
// 		&customer.Created)
// 	if err != nil {
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		}
// 	}

// 	return customer, nil
// }

// type Orders struct {
// 	ID       int       `json:"id,omitempty"`
// 	Customer string    `json:"customer,omitempty"`
// 	Created  time.Time `json:"created,omitempty"`
// }

// type OrdersDetails struct {
// 	ID        int       `json:"id,omitempty"`
// 	ProductID int       `json:"product_id,omitempty"`
// 	Status    string    `json:"status,omitempty"`
// 	Created   time.Time `json:"created,omitempty"`
// }

// type Products struct {
// 	ID          int     `json:"id,omitempty"`
// 	Name        string  `json:"name,omitempty"`
// 	Description string  `json:"description,omitempty"`
// 	Price       float64 `json:"price,omitempty"`
// }
