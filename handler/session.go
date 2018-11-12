package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/OkumuraShintarou/peace/app"
	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/model"
	"github.com/OkumuraShintarou/peace/service"
	"github.com/OkumuraShintarou/peace/util"
)

func SessionSignUp(cc *util.CustomContext) {
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
