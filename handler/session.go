package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/config"
	"github.com/OkumuraShintarou/peace/entity"
	"github.com/OkumuraShintarou/peace/service"
	"github.com/OkumuraShintarou/peace/util"
)

type SessionHandler struct {
	dbm       *gorm.DB
	jwtSecret string
}

func NewSessionHandler(dbm *gorm.DB, config *config.Config) *SessionHandler {
	return &SessionHandler{
		dbm:       dbm,
		jwtSecret: config.JWTSecret,
	}
}

/// MEMO: - UserNameを使用してユーザー登録をする
func (sessionHandler *SessionHandler) SignUp(cc *util.CustomContext) {
	var param entity.SignUpParam

	if stderr := cc.BindJSON(&param); stderr != nil {
		err := apperr.NewError(apperr.BindError, stderr)
		cc.AbortError(400, err)
		return
	}

	userService := service.NewUserService(sessionHandler.dbm)
	user, err := userService.FirstOrCreate(param.UserName)
	if err != nil {
		cc.AbortError(400, err)
		return
	}

	token, err := util.NewJwtToken(user.ID, sessionHandler.jwtSecret)
	if err != nil {
		cc.AbortError(400, err)
		return
	}

	cc.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (sessionHandler *SessionHandler) Me(cc *util.CustomContext) {
	UserID, err := cc.GetAddress()

}
