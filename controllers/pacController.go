package controllers

import (
	"pac-sys/data"
	"pac-sys/models"
	"pac-sys/utils"

	"github.com/gin-gonic/gin"
)

func SavePac(c *gin.Context) {
	var pac models.PacModel
	if err := c.ShouldBind(&pac); err != nil {
		utils.ErrorPanic(err)
	}

	data.Save(pac)
	c.Status(200)
}

func GetPac(c *gin.Context) {
	var pac models.PacModel
	if err := c.ShouldBind(&pac); err != nil {
		utils.ErrorPanic(err)
	}

	c.String(200, data.Get(pac.Key))
}
