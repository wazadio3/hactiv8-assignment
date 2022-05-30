package route

import (
	"assignment2/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controller.CreateOrder)
	router.GET("/orders", controller.GetOrders)
	router.PUT("/orders/:orderId", controller.UpdateOrder)
	router.DELETE("/orders/:orderId", controller.DeleteOrder)

	return router
}
