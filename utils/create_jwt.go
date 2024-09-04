package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("hi")

func CreateJWT(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	})

	token, err := claims.SignedString(secret)

	if err != nil {
		return "", err
	}

	return token, nil
}
