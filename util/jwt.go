package util

import (
	"github.com/OkumuraShintarou/peace/apperr"
	jwt "github.com/dgrijalva/jwt-go"
)

func NewJwtToken(userName, jwtSecret string) (string, *apperr.Error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": userName,
	})

	token, stderr := t.SignedString([]byte(jwtSecret))
	if stderr != nil {
		return "", apperr.NewError(apperr.ServerError, stderr)
	}

	return token, nil
}
