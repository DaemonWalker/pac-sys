package main

import (
	"pac-sys/controllers"
	"pac-sys/data"
	"pac-sys/models"

	"github.com/gin-gonic/gin"
)

func main() {

}

func startGin() {
	data.Migrate()
	r := gin.New()
	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		info := err.(models.PanicInfo)
		c.String(info.Code, info.Message)
	}))
	r.POST("/api/save", controllers.SavePac)
	r.GET("/api/get", controllers.GetPac)
	r.Run("127.0.0.1:5173")
}

func startProxy() {

}
