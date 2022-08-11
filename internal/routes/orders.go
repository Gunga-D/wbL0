package routes

import (
	"encoding/json"
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

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	if err := ordRoute.service.CreateOrder(input); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Status(http.StatusOK)
}

func (ordRoute *OrdersRoute) getOrder(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	result, err := ordRoute.service.GetOrder(id.(string))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	normalizedResult, err := json.Marshal(result)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, normalizedResult)
}
