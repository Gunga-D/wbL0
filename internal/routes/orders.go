package routes

import (
	"fmt"
	"net/http"

	"github.com/Gunga-D/taskL0/internal/models"
	"github.com/Gunga-D/taskL0/internal/services"
	"github.com/gin-gonic/gin"
)

func NewOrdersRoute(service *services.OrderService) *OrdersRoute {
	return &OrdersRoute{service}
}

type OrdersRoute struct {
	service *services.OrderService
}

func (ordRoute *OrdersRoute) addOrder(c *gin.Context) {
	var input models.Order

	err := c.BindJSON(&input)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := ordRoute.service.CreateOrder(input); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (ordRoute *OrdersRoute) getOrder(c *gin.Context) {
	id := c.Param("id")

	result, err := ordRoute.service.GetOrder(id)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
