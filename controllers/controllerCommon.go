package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pac-sys/utils"
)

func BindAction(e *gin.Engine) {
	e.POST("/api/user/getToken", getToken)
}

func bindValue[T any](c *gin.Context) T {
	var t T
	err := c.ShouldBind(&t)
	if err != nil {
		utils.StatusPanic(http.StatusBadRequest)
	}

	return t
}

func json[T any](c *gin.Context, t T) {
	c.JSON(200, t)
}
