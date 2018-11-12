package dao

import (
	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/model"
	"github.com/jinzhu/gorm"
)

const usersTableName = "users"

type User interface {
	FirstOrCreate(user *model.User) (*model.User, *apperr.Error)
}

type UserImpl struct {
	dbm *gorm.DB
}

// ここの戻り値がUserでいける理由はUserImplの構造体がUserのメソッドに準拠しているからである
func NewUser(dbm *gorm.DB) User {
	return &UserImpl{
		dbm: dbm,
	}
}

func (u *UserImpl) FirstOrCreate(newUser *model.User) (*model.User, *apperr.Error) {
	var (
		user model.User,
		err *apperr.Error
	)

	if res := u.dbm.
		Where(model.User{ID: newUser.ID}).
		Attrs(newUser).
		FirstOrCreate(&user); len(res.GetErrors()) > 0 {
			dberr := res.GetErros()[0]
			return nil, apperr.NewError(apperr.DBError, apperr.Warn, dberr.Error())
		}
	return &user, err
}