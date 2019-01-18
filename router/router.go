package router

import (
	"github.com/OkumuraShintarou/peace/app"
	"github.com/OkumuraShintarou/peace/handler"
	"github.com/OkumuraShintarou/peace/util"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	dbm := app.Shared().Dbm
	config := app.Shared().Config

	r := gin.Default()
	r.GET("api/ping", util.CustomHandlerFunc(handler.Ping))

	api := r.Group("api")
	{
		api.POST("/signup_guest", util.CustomHandlerFunc(handler.SessionSignUp))
	}

	return r
}
