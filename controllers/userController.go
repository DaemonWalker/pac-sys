package controllers

import (
	"github.com/gin-gonic/gin"
	"pac-sys/data"
	"pac-sys/models"
	"pac-sys/utils"
)

func getToken(c *gin.Context) {
	account := bindValue[string](c)

	userToken := data.GetUserInfoForLogin(account)
	accessToken := utils.GenerateToken(userToken)

	tokenModel := models.TokenModel{AccessToken: accessToken}
	json(c, tokenModel)
}
