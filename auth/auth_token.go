package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func (a Auth) CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
	})

	tokenString, err := token.SignedString(a.TokenSecret)
	return tokenString, err
}

func (a Auth) GetUserFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.TokenSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user"].(string), nil
	} else {
		return "", err
	}
}
