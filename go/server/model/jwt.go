package model

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	UserID   uint     `json:"userId"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}
