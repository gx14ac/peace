package batch

import (
	"github.com/OkumuraShintarou/peace/client/hackernews"
	"github.com/OkumuraShintarou/peace/entity"
)

func getHackerNews() (*entity.HackerNew, error) {
	hns := hackernews.NewClient()
	hackerNews, err := hns.GET("15723926")
	if err != nil {
		return nil, err
	}

	return hackerNews, nil
}
