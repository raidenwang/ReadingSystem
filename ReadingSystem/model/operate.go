package model

type Collection struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	PublishTime string `json:"publishTime"`
	Link        string `json:"link"`
}
