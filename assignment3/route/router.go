package route

import (
	"assignment3/service"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")

	router.GET("/data", service.GetData)

	return router
}
