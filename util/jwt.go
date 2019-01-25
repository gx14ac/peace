package util

import (
	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/entity"
	jwt "github.com/dgrijalva/jwt-go"
)

func NewJwtToken(user *entity.User, jwtSecret string) (string, *apperr.Error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
	})

	token, stderr := t.SignedString([]byte(jwtSecret))
	if stderr != nil {
		return "", apperr.NewError(apperr.ServerError, stderr)
	}

	return token, nil
}
