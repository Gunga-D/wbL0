package services

import (
	"github.com/Gunga-D/taskL0/internal/models"
	"github.com/Gunga-D/taskL0/internal/repository"
	"github.com/Gunga-D/taskL0/internal/storage"
)

func NewOrderService(cache storage.OrdersMemory, repos repository.OrdersTable) *OrderService {
	return &OrderService{cache, repos}
}

type OrderService struct {
	cache storage.OrdersMemory
	repos repository.OrdersTable
}

func (ord *OrderService) CreateOrder(what models.Order) error {
	ord.cache.Insert(what.OrderUID, what)
	return ord.repos.CreateOrder(what)
}

func (ord *OrderService) GetOrder(key string) (models.Order, error) {
	order, found := ord.cache.Get(key)
	if !found {
		return ord.repos.GetOrder(key)
	}
	return order, nil
}
