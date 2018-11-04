package util

import (
	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/gin-gonic/gin"
)

const (
	errorKey = "error"
)

type CustomContext struct {
	*gin.Context
}

func NewCustomContext(c *gin.Context) *CustomContext {
	return &CustomContext{c}
}

type HandlerFunc func(*CustomContext)

func CustomHandlerFunc(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		cc := &CustomContext{c}
		handler(cc)
	}
}

func (cc *CustomContext) AbortWithError(err *apperr.Error) {
	cc.Set(errorKey, err)
	cc.AbortWithStatusJSON(err.Resp())
}

func (cc *CustomContext) GetError() (*apperr.Error, bool) {
	errIF, ok := cc.Get(errorKey)
	if !ok {
		return nil, false
	}
	err, ok := errIF.(*apperr.Error)
	return err, ok
}
