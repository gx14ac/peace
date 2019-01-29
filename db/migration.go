package db

import (
	"github.com/OkumuraShintarou/peace/entity"
	"github.com/jinzhu/gorm"
)

func Migrate(dbm *gorm.DB) {
	if res := dbm.Debug().AutoMigrate(
		&entity.User{},
	); len(res.GetErrors()) > 0 {
		panic(res.GetErrors()[0])
	}
}
