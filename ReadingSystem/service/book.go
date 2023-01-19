package service

import (
	"go_code/ReadingSystem/dao"
	"go_code/ReadingSystem/model"
)

func GetBookNum() int64 {
	var num int64
	num = dao.GetBookNum()
	return num
}
func GetBookUnLogin(num int64) (data [10]model.BookBasic, err error) {
	data, err = dao.GetBookUnLogin(num)
	return
}
func GetBookLogin(num int64, UserId int64) (data [10]model.BookBasic, err error) {
	data, err = dao.GetBookLogin(num, UserId)
	return
}
func SearchBookByName(BookName string, UserId int64) (u model.BookBasic, err error) {
	u, err = dao.SearchBookByName(BookName, UserId)
	return
}
func StarBook(BookId int, UserId int64) (u model.FavourList, result int) {
	u, result = dao.StarBook(BookId, UserId)
	return
}
func GetBookByLabel(Label string, UserId int64, num int) (data [10]model.BookBasic, err error, n int) {
	data, err, n = dao.GetBookByLabel(Label, UserId, num)
	return
}
