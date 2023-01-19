package service

import (
	"database/sql"
	"go_code/ReadingSystem/dao"
	"go_code/ReadingSystem/model"
)

func GetCommentByBookId(BookId string, UserId int64) (comment [100]model.ApiComment, err error) {
	comment, err = dao.GetCommentByBookId(BookId, UserId)
	return
}
func WriteComment(BookId string, Content string, UserId int64) (result sql.Result, err error, id int64) {
	result, err, id = dao.WriteComment(BookId, Content, UserId)
	return
}
func DeleteComment(CommentId string) (result sql.Result, err error, num1 int64) {
	result, err, num1 = dao.DeleteComment(CommentId)
	return
}
func UpdateComment(CommentId string, Content string) error {
	err := dao.UpdateComment(CommentId, Content)
	return err
}
