package services

import (
	"github.com/Gunga-D/taskL0/internal/models"
	"github.com/Gunga-D/taskL0/internal/repository"
)

func NewOrderService(repos repository.OrdersTable) *OrderService {
	return &OrderService{repos}
}

type OrderService struct {
	repos repository.OrdersTable
}

func (ord *OrderService) CreateOrder(what models.Order) error {
	return ord.repos.CreateOrder(what)
}

func (ord *OrderService) GetOrder(key string) (models.Order, error) {
	return ord.repos.GetOrder(key)
}
