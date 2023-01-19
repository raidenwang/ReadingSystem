package model

type BookBasic struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	IsStar      bool    `json:"isStar"`
	Author      string  `json:"author"`
	CommentNum  int64   `json:"commentNum"`
	Score       float64 `json:"score"`
	Cover       string  `json:"cover"`
	PublishTime string  `json:"publishTime"`
	Link        string  `json:"link"`
	Label       string  `json:"label"`
}
type FavourList struct {
	Id     int64 `json:"id"`
	BookId int64 `json:"bookId"`
	UserId int64 `json:"userId"`
}
