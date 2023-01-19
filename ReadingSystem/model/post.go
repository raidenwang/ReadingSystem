package model

type Post struct {
	Id          int64
	Title       string
	Content     string
	UserId      string
	NickName    string
	Avatar      string
	PublishTime string
	PraiseNum   string
	IsPraised   bool
}
type PostPraisedList struct {
	Id     int64
	PostId int64
	UserId int64
}
