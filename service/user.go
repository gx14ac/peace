package service

import (
	"github.com/jinzhu/gorm"
	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/model"
	"github.com/OkumuraShintarou/peace/dao"
)

type User interface {
	SignUpGuest(param model.SignUpGuestParam) (*model.User, *apperr.Error)
}

type UserImpl struct {
	userDao dao.User
}

func NewUser(dbm *gorm.DB) User {
	return &UserImpl {
		userDao: dao.NewUser(dbm),
	}
}

func (u *UserImpl) SignUpGuest(param model.SignUpGuestParam) (user *model.User, err *apperr.Error) {

	// CreateNewGuestUser
	newUser := model.NewUser()
	newUser.Name = param.Name
	
	// Valodation
	user, err = u.userDao.FirstOrCreate(&newUser)
	return
}