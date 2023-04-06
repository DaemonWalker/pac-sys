package controllers

import (
	"pac-sys/data"
	"pac-sys/models"

	"github.com/gin-gonic/gin"
)

func SavePac(c *gin.Context) {
	var pac models.PacModel
	if c.ShouldBind(&pac) != nil {
		c.Status(400)
		return
	}

	data.Save(pac)
	c.Status(200)
}

func GetPac(c *gin.Context) {
	var pac models.PacModel
	if c.ShouldBind(&pac) == nil {
		c.String(200, data.Get(pac.Key))
		return
	} else {
		c.Status(400)
	}

}
