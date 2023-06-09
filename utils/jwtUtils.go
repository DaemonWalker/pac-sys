package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"pac-sys/constants"
	"pac-sys/entities"
	"pac-sys/share"
	"strings"
	"time"
)

var jwtSecret = "password"
var jwtSecretBytes = []byte(jwtSecret)
var expireTime = 30 * time.Minute
var issuer = "pac-sys"

func GenerateToken(user entities.UserTokenDto) string {
	nowTime := time.Now()
	expireAt := nowTime.Add(expireTime).Unix()

	var claims = entities.ClaimDto{
		Sid:    user.UserId,
		Groups: user.Groups,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecretBytes)

	if err != nil {
		share.CreatePanic(http.StatusInternalServerError, err.Error())
	}

	return token
}

func VerifyToken(tokenString string) (*entities.ClaimDto, error) {
	array := strings.Split(tokenString, "Bearer ")
	if len(array) != 2 {
		return nil, errors.New("invalid token")
	}
	tokenString = array[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecretBytes, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		groupsJson := claims["groups"].([]interface{})
		groups := share.ConvertArray(groupsJson, func(ts any) int {
			return int(ts.(float64))
		})

		claims := entities.ClaimDto{
			Sid:    claims[constants.Sid].(string),
			Groups: groups,
		}
		return &claims, err
	} else {
		return nil, err
	}
}

func SaveAuthorizeInfo(c *gin.Context, claims *entities.ClaimDto) {
	c.Set(constants.AuthKey, claims)
}

func GetAuthorizeInfo(c *gin.Context) *entities.ClaimDto {
	if claims, exists := c.Get(constants.AuthKey); exists {
		return claims.(*entities.ClaimDto)
	} else {
		share.CreatePanic(401, "authorization info is invalid")
		return nil
	}
}
