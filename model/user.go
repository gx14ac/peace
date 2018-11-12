package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID          string             `json:"id" gorm:"primary_key;not null;unique;index;" sql:"type:varchar(36)"`
	Name        string             `json:"name" gorm:"not null;"`
	CreatedAt   time.Time          `json:"-" gorm:"column:createdAt"`
	UpdatedAt   time.Time          `json:"-" gorm:"column:updatedAt"`
}

type SignUpGuestParam struct {
	Name string `json:"name" binding:"required"`
}

func NewUser() User {
	return User{
		ID:        uuid.NewV4().String(),
		Name:      "名無し",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) Resp() map[string]interface{} {
	return map[string]interface{}{
		"id": u.ID,
		"name": u.Name,
		"createdAt": u.CreatedAt,
	}
}