package middleware

import (
	"fmt"

	"github.com/OkumuraShintarou/peace/app"
	"github.com/OkumuraShintarou/peace/util"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cc := util.CustomContext{c}
		fmt.Println("jwt token verifying.....")

		bearer := cc.GetHeader("Authorization")

		id, err := util.GetUserID(bearer, app.Shared().Config.JWTSecret)
		if err != nil {
			cc.AbortWithError(401, err)
			return
		}

		cc.SetID(id)
		cc.Next()
	}
}
