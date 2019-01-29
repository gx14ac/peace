package db

import (
	"fmt"

	"github.com/OkumuraShintarou/peace/entity"
	"github.com/jinzhu/gorm"
)

func Migrate(dbm *gorm.DB) {
	if res := dbm.Debug().AutoMigrate(
		&entity.User{},
	); len(res.GetErrors()) > 0 {
		panic(res.GetErrors()[0])
	}

	if res := dbm.Model(&entity.User{}).AddUniqueIndex("idx_users_id", "id"); len(res.GetErrors()) > 0 {
		fmt.Println(res.GetErrors()[0])
	}
}
