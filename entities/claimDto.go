package entities

import "github.com/dgrijalva/jwt-go"

type ClaimDto struct {
	Sid    string   `json:"sid"`
	Groups []string `json:"groups"`
	jwt.StandardClaims
}
