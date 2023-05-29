package repository

import (
	"github.com/vbelouso/coffeshop/internal/models"
)

type OrderRepository interface {
	GetOrderByID(id int) (models.Order, error)
}
