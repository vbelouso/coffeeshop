package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vbelouso/coffeshop/internal/models"
)

var (
	ErrOrderNotFound = errors.New("FromRepository - order not found")
)

type postgresqlOrderRepository struct {
	connectionPool *pgxpool.Pool
}

func NewPostgreSQLOrderRepository(dsn string) (*postgresqlOrderRepository, error) {
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to parse database configuration: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to establish database connection: %w", err)
	}

	return &postgresqlOrderRepository{pool}, nil
}

const getOrder = `-- name: GetOrder :one
SELECT order_id, customer_id, status, created_at FROM orders
WHERE order_id = $1 LIMIT 1
`

func (p *postgresqlOrderRepository) GetOrderByID(id int) (models.Order, error) {
	row := p.connectionPool.QueryRow(context.Background(), getOrder, id)
	var i models.Order
	err := row.Scan(
		&i.OrderID,
		&i.CustomerID,
		&i.Status,
		&i.CreatedAt,
	)

	return i, err
}
