package repository

import (
	"errors"

	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/entity"
	"github.com/jinzhu/gorm"
)

type User interface {
	FirstOrCreate(userId, userName string) (*entity.User, *apperr.Error)
	FindByUserID(userId string) (*entity.User, *apperr.Error)
}

type UserImpl struct {
	dbm *gorm.DB
}

func NewUserRepository(dbm *gorm.DB) User {
	return &UserImpl{
		dbm: dbm,
	}
}

func (userRepo *UserImpl) FirstOrCreate(userId, userName string) (*entity.User, *apperr.Error) {
	var user entity.User

	res := userRepo.dbm.FirstOrCreate(&user, entity.User{ID: userId, Name: userName})

	errSize := len(res.GetErrors())
	if errSize > 0 {
		err := apperr.NewError(apperr.DBError, res.GetErrors()[0])
		return nil, err
	}

	return &user, nil
}

func (userRepo *UserImpl) FindByUserID(userId string) (*entity.User, *apperr.Error) {
	var user entity.User

	res := userRepo.dbm.Where(entity.User{ID: userId}).Find(&user)

	errSize := len(res.GetErrors())
	if errSize > 0 {
		err := apperr.NewError(apperr.NotFound, errors.New("user not found"))
		return nil, err
	}

	return &user, nil
}
