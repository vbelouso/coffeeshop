package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Customer struct {
	ID       int       `json:"id,omitempty"`
	FullName string    `json:"fullName,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

type Order struct {
	ID       int       `json:"id,omitempty"`
	Customer string    `json:"customer,omitempty"`
	Status   string    `json:"status,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

var conn *pgx.Conn

func openDB(dsn string) error {
	var err error

	conn, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		return err
	}

	return conn.Ping(context.Background())
}

func getCustomerByID(customerID int) (*Customer, error) {
	stmt := `
	SELECT CUSTOMER_ID,
	FULL_NAME,
	CREATED_AT
	FROM PUBLIC.CUSTOMERS
	WHERE CUSTOMER_ID = $1;`

	customer := Customer{}
	err := conn.QueryRow(context.Background(), stmt, customerID).Scan(
		&customer.ID,
		&customer.FullName,
		&customer.Created,
	)

	if err != nil {
		return &Customer{}, fmt.Errorf("failed to get item by ID: %w", err)
	}

	return &customer, nil
}

func getOrderByID(orderID int) (*Order, error) {
	stmt := `
	SELECT ORDER_ID,
	CUSTOMER_ID,
	STATUS,
	CREATED_AT
	FROM PUBLIC.ORDERS
	WHERE ORDER_ID = $1;`

	order := Order{}
	err := conn.QueryRow(context.Background(), stmt, orderID).Scan(
		&order.ID,
		&order.Customer,
		&order.Status,
		&order.Created,
	)

	if err != nil {
		return &Order{}, fmt.Errorf("failed to get item by ID: %w", err)
	}

	return &order, nil
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
