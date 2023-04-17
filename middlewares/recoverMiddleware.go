package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pac-sys/models"
)

func RecoverMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		info, ok := err.(models.PanicInfo)
		if ok {
			c.String(info.Code, info.Message)
		} else {
			c.String(http.StatusInternalServerError, fmt.Sprintf("%v", err))
		}
	})
}
