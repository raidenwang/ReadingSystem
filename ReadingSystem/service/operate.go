package service

import (
	"go_code/ReadingSystem/dao"
	"go_code/ReadingSystem/model"
)

func GetFavourList(UserId int64) (data []model.Collection, err error) {
	data, err = dao.GetFavourList(UserId)
	return
}
func AddFocused(FocusUserId int64, UserId int64) (result int) {
	result = dao.AddFocused(FocusUserId, UserId)
	return
}
func PraiseComment(CommentId string, UserId int64) (result int) {
	result = dao.PraiseComment(CommentId, UserId)
	return
}
func PraisePost(PostId string, UserId int64) (result int) {
	result = dao.PraisePost(PostId, UserId)
	return
}
