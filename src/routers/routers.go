package routers

import (
	"go_simple_mvc/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/test/:name", controllers.GetHello)

	return router
}
