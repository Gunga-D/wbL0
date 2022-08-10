package postgres

import (
	"github.com/Gunga-D/taskL0/internal/models"
	"github.com/jmoiron/sqlx"
)

func NewPostgresOrdersTable(kernel *sqlx.DB) PostgresOrdersTable {
	return PostgresOrdersTable{kernel}
}

type PostgresOrdersTable struct {
	kernel *sqlx.DB
}

func (ordTab PostgresOrdersTable) CreateOrder(what models.Order) error {
	//TODO
	return nil
}

func (ordTab PostgresOrdersTable) GetOrder(key int) (models.Order, error) {
	//TODO
	return models.Order{}, nil
}
