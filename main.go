package main

import (
	"fmt"
	"time"

	"github.com/OkumuraShintarou/peace/app"
	"github.com/OkumuraShintarou/peace/config"
	"github.com/OkumuraShintarou/peace/db"
	resource "github.com/OkumuraShintarou/peace/resources/strings"
	"github.com/OkumuraShintarou/peace/router"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	resource.Peace()
	fmt.Println("Set Time Zone...", time.UTC)
	time.Local = time.UTC
}

func main() {
	cfg := config.New()
	dbm := db.MustNewDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	app.Init(dbm, cfg)
	db.Migrate(dbm)
	r := router.New()
	r.Run(":" + cfg.AppPort)
}
