package handler

import (
	"github.com/OkumuraShintarou/peace/app"
	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/OkumuraShintarou/peace/model"
	"github.com/OkumuraShintarou/peace/service"
	"github.com/OkumuraShintarou/peace/util"
)

func SessionSignUp(cc *util.CustomContext) {
	var param model.SignUpParam

	if bindErr := cc.BindJSON(&param); bindErr != nil {
		err := apperr.NewError(apperr.RequestError, apperr.Info, bindErr.Error())
		cc.AbortWithError(err)
		return
	}

	userSvc := service.NewUser(app.DBM())
}
