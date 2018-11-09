package main

import (
	"fmt"
	"time"

	"github.com/OkumuraShintarou/peace/router"
	"github.com/OkumuraShintarou/peace/app"
	"github.com/OkumuraShintarou/peace/config"
	"github.com/OkumuraShintarou/peace/db"
	resource "github.com/OkumuraShintarou/peace/resources/strings"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.New()
	dbm := db.MustNewDB(cfg.DBPass, cfg.DBName)
	app.Init(dbm, cfg)

	r := router.New()
	r.Run(":" + cfg.AppPort)
}

func init() {
	resource.Peace()
	fmt.Println("Set Time Zone...", time.UTC)
	time.Local = time.UTC
}
