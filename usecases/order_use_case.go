package usecases

import (
	"boon/entities"
	"errors"
)

type OrderUseCase interface {
	CreateOrder(order entities.Order) error
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderUseCase {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order entities.Order) error {

	if order.Total <= 0 {
		return errors.New("total must be more than 0")
	}
	return s.repo.Save(order)
}
