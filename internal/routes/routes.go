package routes

import "github.com/gin-gonic/gin"

func Init(ordRoute *OrdersRoute) *gin.Engine {
	router := gin.New()

	api := router.Group("/orders")
	{
		api.POST("/", ordRoute.addOrder)
		api.GET("/:id", ordRoute.getOrder)
	}

	return router
}
