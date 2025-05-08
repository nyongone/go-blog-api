package util

import (
	"errors"
	"go-blog-api/domain"
	"go-blog-api/ent"
	"go-blog-api/internal/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(user *ent.User) (string, error)  {
	JWTATHour, err := strconv.Atoi(config.EnvVar.JWTATHour)

	if err != nil {
		return "", err
	}

	exp := time.Now().Add(time.Hour * time.Duration(JWTATHour)).Unix()
	claims := &domain.JWTAccessTokenClaims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	at, err := token.SignedString([]byte(config.EnvVar.JWTSecret))
	if err != nil {
		return "", err
	}

	return at, nil
}

func CreateRefreshToken() (string, error) {
	JWTRTHour, err := strconv.Atoi(config.EnvVar.JWTRTHour)

	if err != nil {
		return "", err
	}

	exp := time.Now().Add(time.Hour * time.Duration(JWTRTHour)).Unix()
	claims := &domain.JWTRefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	rt, err := token.SignedString([]byte(config.EnvVar.JWTSecret))
	if err != nil {
		return "", err
	}

	return rt, nil
}

func VerifyToken(token string) (bool, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.EnvVar.JWTSecret), nil
	})

	if err != nil {
		return false, err
	}

	if !parsedToken.Valid {
		return false, errors.New("invalid jwt token")
	}

	return true, nil
}