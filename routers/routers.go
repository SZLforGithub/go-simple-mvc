package routers

import (
	"github.com/SZLforGithub/go-simple-mvc/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter is used to init router in main.go
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", controllers.GetHello)

	return router
}
