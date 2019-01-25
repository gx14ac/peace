package entity

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primary_key;unique;index;"`
	Name      string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}
