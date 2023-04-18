package entities

import (
	"github.com/dgrijalva/jwt-go"
	"pac-sys/constants"
	"pac-sys/share"
)

type ClaimDto struct {
	Sid    string `json:"sid"`
	Groups []int  `json:"groups"`
	jwt.StandardClaims
}

func (c ClaimDto) IsAdmin() bool {
	return share.ArrayContains(c.Groups, constants.AdminGroupId)
}
