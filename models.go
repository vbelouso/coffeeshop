package main

import (
	"context"
	"fmt"
	"os"
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
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
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

	customer := &Customer{}
	err := conn.QueryRow(context.Background(), stmt, customerID).Scan(
		&customer.ID,
		&customer.FullName,
		&customer.Created,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get item by ID: %w", err)
	}

	return customer, nil
}

func getCustomers() ([]*Customer, error) {
	stmt := `
	SELECT CUSTOMER_ID,
	FULL_NAME,
	CREATED_AT
	FROM PUBLIC.CUSTOMERS;`

	rows, err := conn.Query(context.Background(), stmt)
	if err != nil {
		return nil, fmt.Errorf("failed to get customers: %w", err)
	}

	defer rows.Close()

	var customers []*Customer
	for rows.Next() {
		c := &Customer{}
		err = rows.Scan(
			&c.ID,
			&c.FullName,
			&c.Created,
		)
		if err != nil {
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func getOrderByID(orderID int) (*Order, error) {
	stmt := `
	SELECT ORDER_ID,
	CUSTOMER_ID,
	STATUS,
	CREATED_AT
	FROM PUBLIC.ORDERS
	WHERE ORDER_ID = $1;`

	order := &Order{}
	err := conn.QueryRow(context.Background(), stmt, orderID).Scan(
		&order.ID,
		&order.Customer,
		&order.Status,
		&order.Created,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get item by ID: %w", err)
	}

	return order, nil
}

func getOrders() ([]*Order, error) {
	stmt := `
	SELECT ORDER_ID,
	CUSTOMER_ID,
	STATUS,
	CREATED_AT
	FROM PUBLIC.ORDERS;`

	rows, err := conn.Query(context.Background(), stmt)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}

	defer rows.Close()

	var orders []*Order

	for rows.Next() {
		o := &Order{}
		err = rows.Scan(&o.ID, &o.Customer, &o.Status, &o.Created)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
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
