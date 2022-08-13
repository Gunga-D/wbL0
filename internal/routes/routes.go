package routes

import "github.com/gin-gonic/gin"

func Create(ordRoute *OrdersRoute) *gin.Engine {
	engine := gin.New()

	engine.LoadHTMLGlob("./static/*.html")

	api := engine.Group("/api")
	{
		orders := api.Group("/orders")
		{
			orders.POST("/", ordRoute.addOrder)
			orders.GET("/:id", ordRoute.getOrder)
		}
	}

	view := engine.Group("/view")
	{
		orders := view.Group("/orders")
		{
			orders.GET("/:id", ordRoute.renderOrder)
		}
	}

	return engine
}
