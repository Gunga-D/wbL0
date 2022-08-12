package postgres

import (
	"encoding/json"
	"fmt"

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
	normalizedOrder, err := json.Marshal(what)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO orders (order_uid, data) VALUES ($1, $2)")
	ordTab.kernel.QueryRow(query, what.OrderUID, normalizedOrder)

	return nil
}

func (ordTab PostgresOrdersTable) GetOrder(key string) (models.Order, error) {
	//TODO
	return models.Order{}, nil
}