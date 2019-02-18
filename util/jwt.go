package util

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/entity"
	jwt "github.com/dgrijalva/jwt-go"
)

/// MEMO: - Tokenを作成する
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

// Userからリクエストが来た時にtokenからuserのidをSetする
func GetUserID(bearer, jwtSecret string) (string, *apperr.Error) {
	if strings.Contains(bearer, "Bearer") {
		bearer = strings.Replace(bearer, "Bearer", "", 1)
		if bearer == "" {
			return "", apperr.NewError(apperr.AuthorizationError, errors.New("no authroization"))
		}
	}

	token, stderr := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})
	if stderr != nil {
		return "", apperr.NewError(apperr.AuthorizationError, stderr)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", apperr.NewError(apperr.AuthorizationError, errors.New("token invalid"))
	}

	id, ok := claims["id"].(string)
	if !ok {
		return "", apperr.NewError(apperr.AuthorizationError, errors.New("id cast error"))
	}

	return id, nil
}
