package app

import (
	"github.com/OkumuraShintarou/peace/config"
	"github.com/OkumuraShintarou/peace/db"
	"github.com/jinzhu/gorm"
)

type App struct {
	Dbm    *gorm.DB
	Config config.Config
}

var sharedInstance *App

func init() {
	cfg := config.NewConfig()
	dbm := db.MustNewDB(cfg.DBHost, cfg.AppPort, cfg.DBUser, cfg.DBName)

	sharedInstance = &App{
		Dbm:    dbm,
		Config: cfg,
	}
}

func Shared() *App {
	return sharedInstance
}
