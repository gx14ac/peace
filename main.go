package main

import (
	"fmt"
	"time"

	"github.com/OkumuraShintarou/peace/app"
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

	appPort := cfg.AppPort
	db.Migrate(dbm)
	r := router.NewRouter()
	r.Run(":" + appPort)
}
