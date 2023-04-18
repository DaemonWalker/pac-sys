package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pac-sys/data"
	"pac-sys/entities"
	"pac-sys/models"
	"pac-sys/share"
	"pac-sys/utils"
)

func getAllPacs(c *gin.Context) {
	claims := utils.GetAuthorizeInfo(c)
	pacs := data.GetPacByGroupId(claims.Groups)

	json(c, pacs)
}

func savePac(c *gin.Context) {
	claims := utils.GetAuthorizeInfo(c)
	pac := bindValue[models.CreatePacModel](c)

	if !claims.IsAdmin() && !share.ArrayContains(claims.Groups, pac.GroupId) {
		c.AbortWithError(http.StatusBadRequest, errors.New("you cannot create pac with group that you are not belong to"))
	}

	data.SavePac(entities.PacEntity{GroupId: pac.GroupId, Name: pac.Name, Value: pac.Value})
}
