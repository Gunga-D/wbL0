package storage

import (
	"github.com/Gunga-D/taskL0/internal/models"
	"github.com/Gunga-D/taskL0/internal/repository/postgres"
)

func NewOrdersMemory() OrdersMemory {
	return OrdersMemory{make(map[string]models.Order)}
}

type OrdersMemory struct {
	core map[string]models.Order
}

func (m *OrdersMemory) Recover(from *postgres.PostgresOrdersTable) error {
	models, err := from.GetAll()
	if err != nil {
		return err
	}

	for _, model := range models {
		m.core[model.OrderUID] = model
	}
	return nil
}

func (m *OrdersMemory) Get(key string) (models.Order, bool) {
	result, found := m.core[key]
	return result, found
}

func (m *OrdersMemory) Insert(key string, order models.Order) {
	m.core[key] = order
}
