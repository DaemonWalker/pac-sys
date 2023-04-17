package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pac-sys/utils"
)

var anonymousPath []string = []string{"/api/user/getToken"}

const AuthorizationHeader = "Authorization"

func AuthMiddleware(c *gin.Context) {
	anonymous := false

	for _, path := range anonymousPath {
		if path == c.Request.RequestURI {
			anonymous = true
		}
	}
	if anonymous {
		return
	}

	tokenString := c.Request.Header.Get(AuthorizationHeader)
	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	utils.SaveAuthorizeInfo(c, claims)
}
