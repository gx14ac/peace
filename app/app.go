package app

import (
	"github.com/OkumuraShintarou/peace/config"
	"github.com/jinzhu/gorm"
)

type App struct {
	dbm    *gorm.DB
	config config.Config
}

var _app *App

func Init(dbm *gorm.DB, cfg config.Config) {
	_app = &App{
		dbm:    dbm,
		config: cfg,
	}
}

func DBM() *gorm.DB {
	return _app.dbm
}

func IsPrd() bool {
	return _app.config.AppEnv == "prd"
}

func IsDev() bool {
	return _app.config.AppEnv == "dev"
}
