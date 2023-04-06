package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"pac-sys/controllers"
	"pac-sys/data"
)

func main() {
	data.Migrate()
	r := gin.Default()
	r.POST("/api/save", controllers.SavePac)
	r.GET("/api/get", controllers.GetPac)
	r.Run() // listen and serve on 0.0.0.0:8080
	fmt.Println("asd")
}
