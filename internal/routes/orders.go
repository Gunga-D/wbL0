package routes

import (
	"github.com/Gunga-D/taskL0/internal/services"
	"github.com/gin-gonic/gin"
)

func NewOrdersRoute(service *services.OrderService) *OrdersRoute {
	return &OrdersRoute{service}
}

type OrdersRoute struct {
	service *services.OrderService
}

func (ordRoute *OrdersRoute) getOrder(c *gin.Context) {
	//TODO
}

func (ordRoute *OrdersRoute) addOrder(c *gin.Context) {
	//TODO
}
