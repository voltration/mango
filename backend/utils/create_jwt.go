package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Secret = []byte("hi")

func CreateJWT(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	})

	token, err := claims.SignedString(Secret)

	if err != nil {
		return "", err
	}

	return token, nil
}
