package handler

import (
	"net/http"

	"github.com/OkumuraShintarou/peace/util"
)

func Ping(cc *util.CustomContext) {
	cc.String(http.StatusOK, "pong")
}
