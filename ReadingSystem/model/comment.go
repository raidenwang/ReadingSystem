package model

type Comment struct {
	Id          int64  `json:"id"`
	BookId      string `json:"bookId"`
	PublishTime string `json:"publishTime"`
	Content     string `json:"content"`
	UserId      int64  `json:"userId"`
	Avatar      string `json:"avatar"`
	NickName    string `json:"nickName"`
	PraiseCount int    `json:"praiseCount"`
	IsPraised   bool   `json:"isPraised"`
	IsFocus     bool   `json:"isFocus"`
}
type PostComment struct {
	Content string `json:"content"`
}
type PostCommentID struct {
	CommentID string `json:"commentID"`
}
type ApiComment struct {
	BookId      string `json:"bookId"`
	PublishTime string `json:"publishTime"`
	Content     string `json:"content"`
	UserId      int64  `json:"userId"`
	Avatar      string `json:"avatar"`
	NickName    string `json:"nickName"`
	PraiseCount int    `json:"praiseCount"`
	IsPraised   bool   `json:"isPraised"`
	IsFocus     bool   `json:"isFocus"`
}
type PraisedList struct {
	Id        int64 `json:"id"`
	CommentId int64 `json:"commentId"`
	UserId    int64 `json:"userId"`
}
type FocusedList struct {
	Id            int64 `json:"id"`
	UserId        int64 `json:"userId"`
	FocusedUserId int64 `json:"focusedUserId"`
}
