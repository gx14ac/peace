package db

import (
	"github.com/jinzhu/gorm"
	"github.com/OkumuraShintarou/peace/model"
)

func Migrate(dbm *gorm.DB) {
	if res := dbm.Debug().AutoMigrate(
		&model.User{},
	); len(res.GetErrors()) > 0 {
		panic(res.GetErrors()[0])
	}
}