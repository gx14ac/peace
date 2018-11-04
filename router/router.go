package router

import (
	"github.com/OkumuraShintarou/peace/util"
	"github.com/OkumuraShintarou/peace/handler"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("api/ping", util.CustomHandlerFunc(handler.Ping))
	return r
}
