package service

import (
	"github.com/jinzhu/gorm"

	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/entity"
	"github.com/OkumuraShintarou/peace/repository"
)

type UserService struct {
	userRepo repository.User
}

func NewUserService(dbm *gorm.DB) *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(dbm),
	}
}

func (userService *UserService) FirstOrCreate(userName string) (*entity.User, *apperr.Error) {
	return userService.FirstOrCreate(userName)
}
