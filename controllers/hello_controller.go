package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	c.String(http.StatusOK, "Hello %s!", name)
}
