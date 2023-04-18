package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pac-sys/models"
)

func RecoverMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		info, ok := err.(models.PanicInfo)
		if ok {
			c.AbortWithError(info.Code, errors.New(info.Message))
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	})
}
