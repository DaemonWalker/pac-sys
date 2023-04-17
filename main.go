package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"pac-sys/controllers"
	"pac-sys/data"
	. "pac-sys/middlewares"
)

func main() {
	data.InitDB()

	r := gin.New()
	r.Use(
		RecoverMiddleware(),
		AuthMiddleware,
		gin.Logger(),
	)

	controllers.BindAction(r)

	err := r.Run("127.0.0.1:5173")
	if err != nil {
		log.Fatalln(err)
	}
}
