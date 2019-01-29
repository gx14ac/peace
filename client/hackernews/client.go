package hackernews

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/OkumuraShintarou/peace/entity"
)

type client struct {
	httpClient    *http.Client
	defaultHeader http.Header
}

type HackerNewsClientService interface {
	GET(id string) (*entity.HackerNew, error)
}

type HackerNewsClientServiceImpl struct {
	client client
}

func NewClient() HackerNewsClientService {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			// InsecureSkipVerify: true, 中間車攻撃をうける可能性があるのでフラグは指定しない
		},
	}

	hackerNewsImpl := HackerNewsClientServiceImpl{
		client: client{
			httpClient:    &http.Client{Transport: tr},
			defaultHeader: make(http.Header),
		},
	}

	hackerNewsImpl.client.defaultHeader.Set("Content-Type", "application/json")
	hackerNewsImpl.client.defaultHeader.Set("Cache-Control", "no-cache")

	return hackerNewsImpl
}

func (hnc HackerNewsClientServiceImpl) GET(id string) (*entity.HackerNew, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%s.json",
		id,
	)
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := hnc.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type responseBody struct {
		OwnerName   string `json:"by"`
		Descendants int64  `json:"descendants"`
		NewsID      int64  `json:"id"`
		Kids        []int  `json:"kids"`
		Score       int64  `json:"score"`
		Time        int64  `json:"time"`
		Title       string `json:"title"`
		Type        string `json:"type"`
		URL         string `json:"url"`
	}

	var respBody responseBody
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return nil, err
	}

	hn := &entity.HackerNew{
		OwnerName:   respBody.OwnerName,
		Descendants: respBody.Descendants,
		NewsID:      respBody.NewsID,
		Kids:        respBody.Kids,
		Score:       respBody.Score,
		Time:        respBody.Time,
		Title:       respBody.Title,
		Type:        respBody.Type,
		URL:         respBody.URL,
	}

	fmt.Println(hn)
	return hn, nil
}
