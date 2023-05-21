package service

import (
	"github.com/vbelouso/coffeshop/internal/models"
)

type IOrderService interface {
	GetOrderByID(id int) (models.Order, error)
}
