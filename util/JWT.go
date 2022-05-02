package util

import (
	"github.com/golang-jwt/jwt"
)

func VerifyToken(token string) bool {

	_, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetEnv("JWT_SECRET", "secret")), nil
	})

	if err != nil {
		return false
	}

	return true
}

func GetToken(token string) *jwt.Token {

	t, _ := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetEnv("JWT_SECRET", "secret")), nil
	})

	return t
}
