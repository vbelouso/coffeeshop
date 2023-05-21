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

type OrderService struct {
	orderRepo repository.IOrderRepository
}

func NewOrderService(orderRepo repository.IOrderRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (d *OrderService) GetOrderByID(id int) (models.Order, error) {
	if id <= 0 {
		return models.Order{}, ErrIDIsNotValid
	}

	order, err := d.orderRepo.GetOrderByID(id)
	if errors.Is(err, repository.ErrOrderNotFound) {
		return models.Order{}, ErrOrderNotFound
	}

	return order, nil
}
