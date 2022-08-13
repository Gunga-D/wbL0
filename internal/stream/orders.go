package stream

import (
	"encoding/json"

	"github.com/Gunga-D/taskL0/internal/models"
	"github.com/Gunga-D/taskL0/internal/services"
	"github.com/nats-io/stan.go"
)

func NewOrdersStream(service *services.OrderService) *OrdersStream {
	return &OrdersStream{service}
}

type OrdersStream struct {
	service *services.OrderService
}

func (ordStream *OrdersStream) addOrder(m *stan.Msg) {
	var input models.Order
	if err := json.Unmarshal(m.Data, &input); err != nil {
		return
	}

	ordStream.service.CreateOrder(input)
}
