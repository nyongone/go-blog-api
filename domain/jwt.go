package domain

import "github.com/golang-jwt/jwt"

type JWTAccessTokenClaims struct {
	Email	string	`json:"name"`
	jwt.StandardClaims
}

type JWTRefreshTokenClaims struct {
	jwt.StandardClaims
}