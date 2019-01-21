package util

import (
	"errors"

	"github.com/OkumuraShintarou/peace/apperr"
	"github.com/gin-gonic/gin"
)

const (
	nameKey    = "name"
	addressKey = "address"
	errorKey   = "error"
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

func (cc *CustomContext) AbortError(status int, err *apperr.Error) {
	cc.SetError(err)
	cc.AbortWithStatusJSON(status, err.ToJSON())
}

func (cc *CustomContext) SetError(err *apperr.Error) {
	cc.Set(errorKey, err)
}

func (cc *CustomContext) GetError() (*apperr.Error, bool) {
	errIF, ok := cc.Get(errorKey)
	if !ok {
		return nil, false
	}
	err, ok := errIF.(*apperr.Error)
	return err, ok
}

func (cc *CustomContext) GetUserName() (string, *apperr.Error) {
	addIf, ok := cc.Get(nameKey)
	if !ok {
		return "", apperr.NewError(apperr.ServerError, errors.New("no name found in context"))
	}

	addr, ok := addIf.(string)
	if !ok {
		return "", apperr.NewError(apperr.ServerError, errors.New("name cast error"))
	}

	return addr, nil
}
