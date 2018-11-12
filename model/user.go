package model

import (
	"time"

	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/util"
	"github.com/satori/go.uuid"
)

type User struct {
	ID string
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}