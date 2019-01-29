package entity

type HackerNew struct {
	ID          uint   `json:"-" gorm:"primary_key;unique;index;not null;"`
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
