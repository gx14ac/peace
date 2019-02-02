package batch

import (
	"github.com/OkumuraShintarou/peace/client/hackernews"
	"github.com/OkumuraShintarou/peace/entity"
)

type HackerNewsBatch struct {
	hackerNewsCli hackernews.HackerNewsClientService
}

func NewHackerNewsBatch() *HackerNewsBatch {
	return &HackerNewsBatch{
		hackerNewsCli: hackernews.NewHackerNewsClient(),
	}
}

func (hns HackerNewsBatch) getHackerNews() (*entity.HackerNew, error) {
	hackerNews, err := hns.hackerNewsCli.GET("15723926")
	if err != nil {
		return nil, err
	}

	return hackerNews, nil
}

func (hns HackerNewsBatch) getHackerNewsID() {
	hackerNewsID, err := hns.hackerNewsCli.GETNewsID()
	if err != nil {
		return nil, err
	}
}
