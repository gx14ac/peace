package main

import (
	"fmt"
	"time"

	"github.com/OkumuraShintarou/peace/app"
	"github.com/OkumuraShintarou/peace/batch"
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
	cfg := app.Shared().Config
	dbm := app.Shared().Dbm

	appService := cfg.AppService
	appPort := cfg.AppPort

	db.Migrate(dbm)
	if appService == "api" {
		r := router.NewRouter()
		r.Run(":" + appPort)
	} else if appService == "batch" {
		batch.Start(dbm, cfg)
	}
}
