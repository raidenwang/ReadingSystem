package service

import (
	"database/sql"
	"go_code/ReadingSystem/dao"
	"go_code/ReadingSystem/model"
)

func SearchUserByUserName(name string) (u model.UserBasic, err error) {
	u, err = dao.SearchUserByName(name)
	return
}
func Login(username string, password string) (code int) {
	code = dao.Login(username, password)
	return
}
func CreateUser(name string, password string) (sql.Result, error) {
	result, err := dao.InsertUser(name, password)
	return result, err
}
func GetIdByName(name string) (id int64, err error) {
	id, err = dao.GetIdByName(name)
	return
}
func GetPwdById(UserId int64) (pwd string, err error) {
	pwd, err = dao.GetPwdById(UserId)
	return
}
func SearchUserById(UserId int64) (u model.UserBasic, err error) {
	u, err = dao.SearchUserById(UserId)
	return
}