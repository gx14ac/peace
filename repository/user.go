package repository

import (
	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/entity"
	"github.com/jinzhu/gorm"
)

type User interface {
	FirstOrCreate(userId string) (*entity.User, *apperr.Error)
}

type UserImpl struct {
	dbm *gorm.DB
}

func NewUserRepository(dbm *gorm.DB) User {
	return &UserImpl{
		dbm: dbm,
	}
}

func (userRepo *UserImpl) FirstOrCreate(userId string) (*entity.User, *apperr.Error) {
	var user entity.User

	res := userRepo.dbm.FirstOrCreate(&user, entity.User{ID: userId})

	errSize := len(res.GetErrors())
	if errSize > 0 {
		err := apperr.NewError(apperr.DBError, res.GetErrors()[0])
		return nil, err
	}

	return &user, nil
}