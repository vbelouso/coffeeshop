package repository

import (
	"github.com/vbelouso/coffeshop/internal/models"
)

type IOrderRepository interface {
	GetOrderByID(id int) (models.Order, error)
}
