package service

import (
	"github.com/jinzhu/gorm"

	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/entity"
	"github.com/OkumuraShintarou/peace/helper"
	"github.com/OkumuraShintarou/peace/repository"
)

type UserService struct {
	userRepo repository.User
}

// ここでgormを引数に取らずにwireで実現する
func NewUserService(dbm *gorm.DB) *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(dbm),
	}
}

func (userService *UserService) FirstOrCreate(userName string) (*entity.User, *apperr.Error) {
	userID := helper.CreateUserID()
	return userService.userRepo.FirstOrCreate(userID, userName)
}

func (userService *UserService) FindByUserID(userID string) (*entity.User, *apperr.Error) {
	// helperを使用してuserIdを作成し、userRepoの引数にいれる
	return userService.userRepo.FindByUserID(userID)
}
