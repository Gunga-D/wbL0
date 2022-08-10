package repository

import "github.com/Gunga-D/taskL0/internal/models"

type OrdersTable interface {
	CreateOrder(what models.Order) error
	GetOrder(key int) (models.Order, error)
}
