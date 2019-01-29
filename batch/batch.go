package batch

import (
	"github.com/OkumuraShintarou/peace/config"
	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"
)

func Start(dbm *gorm.DB, cfg *config.Config) {
	gocron.Every(10).Seconds().Do(outputHelloWorld)
	gocron.Every(10).Seconds().Do(getHackerNews)
}
