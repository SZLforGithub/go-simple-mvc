package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHello(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{gin.Param{Key: "name", Value: "Louis"}}

	GetHello(c)

	assert.Equal(t, http.StatusOK, w.Code)

	b, err := io.ReadAll(w.Body)

	assert.Nil(t, err)
	assert.Equal(t, string(b), "Hello Louis!")
}
