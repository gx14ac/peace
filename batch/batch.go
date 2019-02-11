package batch

import (
	"github.com/OkumuraShintarou/peace/config"
	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"
)

func Start(dbm *gorm.DB, cfg *config.Config) {
	nhnb := NewHackerNewsBatch()
	gocron.Every(5).Seconds().Do(nhnb.getHackerNews)
	gocron.Every(5).Seconds().Do(nhnb.getHackerNewsID)

	<-gocron.Start()
}
