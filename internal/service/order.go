package service

import (
	"errors"

	"github.com/vbelouso/coffeshop/internal/models"
	"github.com/vbelouso/coffeshop/internal/repository"
)

var (
	ErrIDIsNotValid  = errors.New("id is not valid")
	ErrOrderNotFound = errors.New("the order cannot be found")
)

type OrderService interface {
	GetOrderByID(id int) (models.Order, error)
}
type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) *orderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (d *orderService) GetOrderByID(id int) (models.Order, error) {
	if id <= 0 {
		return models.Order{}, ErrIDIsNotValid
	}

	order, err := d.orderRepo.GetOrderByID(id)
	if errors.Is(err, repository.ErrOrderNotFound) {
		return models.Order{}, ErrOrderNotFound
	}

	return order, nil
}
