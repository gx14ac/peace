package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string    `json:"id" gorm:"primary_key;not null;unique;index;" sql:"type:varchar(36)"`
	Name      string    `json:"name" gorm:"not null";`
	CreatedAt time.Time `json:"-" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"-" gorm:"column:createdAt"`
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
