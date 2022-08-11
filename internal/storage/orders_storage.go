package storage

import "github.com/Gunga-D/taskL0/internal/models"

type OrdersMemory struct {
	core map[string]models.Order
}

func (m *OrdersMemory) Get(key string) models.Order {
	return m.core[key]
}

func (m *OrdersMemory) Insert(key string, order models.Order) {
	m.core[key] = order
}
