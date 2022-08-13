package postgres

import (
	"encoding/json"

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

	ordTab.kernel.QueryRow("INSERT INTO orders (order_uid, order_data) VALUES ($1, $2)", what.OrderUID, normalizedOrder)

	return nil
}

func (ordTab PostgresOrdersTable) GetOrder(key string) (models.Order, error) {
	dbdata := json.RawMessage{}
	if err := ordTab.kernel.Get(&dbdata, "SELECT order_data FROM orders WHERE order_uid = $1 LIMIT 1", key); err != nil {
		return models.Order{}, err
	}

	result := models.Order{}
	if err := json.Unmarshal(dbdata, &result); err != nil {
		return models.Order{}, err
	}

	return result, nil
}
