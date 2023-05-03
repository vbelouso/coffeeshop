-- name: GetCustomer :one
SELECT * FROM customers
WHERE customer_id = $1 LIMIT 1;

-- name: ListCustomers :many
SELECT * FROM customers
ORDER BY customer_id;

-- name: GetOrder :one
SELECT * FROM orders
WHERE order_id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
ORDER BY order_id;