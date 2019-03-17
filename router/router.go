package router

import (
	"github.com/OkumuraShintarou/peace/app"
	"github.com/OkumuraShintarou/peace/handler"
	"github.com/OkumuraShintarou/peace/middleware"
	"github.com/OkumuraShintarou/peace/util"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	dbm := app.Shared().Dbm
	config := app.Shared().Config

	sessionHandler := handler.NewSessionHandler(dbm, config)

	r := gin.Default()
	// r.GET("api/ping", util.CustomHandlerFunc(handler.Ping))

	api := r.Group("api")
	{
		noAuth := api.Group("")
		api.GET("/ping", util.CustomHandlerFunc(handler.Ping))
		api.POST("/signup_guest", util.CustomHandlerFunc(sessionHandler.SignUp))
		// using Auth API
		jwtAuth := noAuth.Group("")
		jwtAuth.Use(middleware.JwtAuth())
		{
			jwtAuth.GET("/me", util.CustomHandlerFunc(sessionHandler.Me))
			jwtAuth.POST("/me/update", util.CustomHandlerFunc(sessionHandler.Update))
			// jwtAuth.POST("/me/reflresh", util.CustomHandlerFunc())
		}
	}

	return r
}
