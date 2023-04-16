package middlewares

import (
	"github.com/gin-gonic/gin"
	"pac-sys/models"
)

func RecoverMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		info := err.(models.PanicInfo)
		c.String(info.Code, info.Message)
	})
}
