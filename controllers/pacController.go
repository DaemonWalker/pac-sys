package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"pac-sys/utils"
)

func getAllPacs(c *gin.Context) {
	claims := utils.GetAuthorizeInfo(c)
	log.Default().Println("getAllPacs", claims.Sid, claims.Groups)
}
