package routes

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	router := gin.New()

	router.GET("/")
	api := router.Group("/orders")
	{
		api.POST("/", addOrder)
		api.GET("/:id", getOrder)
	}

	return router
}
