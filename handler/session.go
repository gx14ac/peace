package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/OkumuraShintarou/peace/app"
	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/config"
	"github.com/OkumuraShintarou/peace/model"
	"github.com/OkumuraShintarou/peace/service"
	"github.com/OkumuraShintarou/peace/util"
	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	dbm       *gorm.DB
	jwtSecret string
}

func NewSessionHandler(dbm *gorm.DB, config *config.Config) *SessionHandler {
	return &SessionHandler{
		dbm:       dbm,
		jwtSecret: config.JwtSecret,
	}
}

func (sessionHandler *SessionHandler) Sign(cc *util.CustomContext) {

}

// will deprecated
// iOS側からじゃなくても叩けるように作り直す
/// TODO:JwtSecretを使って、ユーザー作成を行う。
func (sessionHandler *SessionHandler) SessionSignUp(cc *util.CustomContext) {
	var param model.SignUpGuestParam

	if bindErr := cc.BindJSON(&param); bindErr != nil {
		err := apperr.NewError(apperr.RequestError, apperr.Info, bindErr.Error())
		cc.AbortWithError(err)
		return
	}

	userSvc := service.NewUser(app.DBM())

	user, err := userSvc.SignUpGuest(param)

	if err != nil {
		cc.AbortWithError(err)
		return
	}

	cc.JSON(http.StatusOK, gin.H{
		"user": user.Resp(),
	})
}
