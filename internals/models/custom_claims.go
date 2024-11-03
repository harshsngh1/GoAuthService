package models

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	Email  string `json:"email"`
	UserId string `json:"userId"`
	jwt.StandardClaims
}
