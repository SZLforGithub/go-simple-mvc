package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s!", name)
	// c.JSON(200, gin.H{"message": "Hello, Gin!"})
}
