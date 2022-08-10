package services

import "github.com/Gunga-D/taskL0/internal/repository"

func NewOrderService(repos repository.OrdersTable) *OrderService {
	return &OrderService{repos}
}

type OrderService struct {
	repos repository.OrdersTable
}

func (ord *OrderService) Create() {
	//TODO
}

func (ord *OrderService) Get() {
	//TODO
}
