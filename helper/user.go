package helper

import (
	uuid "github.com/satori/go.uuid"
)

func CreateUserID() string {
	userID := uuid.NewV4().String()
	return userID
}
