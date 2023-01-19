package dao

import (
	"database/sql"
	"go_code/ReadingSystem/model"
)

func InsertUser(name string, password string) (result sql.Result, err error) {
	result, err = DB.Exec("INSERT into user_basics(name,password)values (?,?)", name, password)
	return result, err
}
func Login(username string, password string) (code int) {
	u, err := SearchUserByName(username)
	if err != nil {
		return
	}
	User00.ID = u.ID
	if u.ID == 0 {
		code = 1
		return
	}
	if password != u.Password {
		code = 2
		return
	}
	code = 3
	return

}
func SearchUserByName(name string) (u model.UserBasic, err error) {
	row := DB.QueryRow("select id,name,password from user_basics where name=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.Name, &u.Password)
	return
}
func SearchUserById(UserId int64) (u model.UserBasic, err error) {
	str := "select id,name,password,nick_name,avatar,introduction,telephone,qq,gender,email,birthday from user_basics where id=?"
	row := DB.QueryRow(str, UserId)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.Name, &u.Password, &u.NickName, &u.Avatar, &u.Introduction, &u.Telephone, &u.QQ, &u.Gender, &u.Email, &u.Birthday)
	return
}
func GetIdByName(name string) (id int64, err error) {
	row := DB.QueryRow(" select id from user_basics where name=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	u0 := new(model.UserBasic)
	err = row.Scan(&u0.ID)
	return u0.ID, nil
}
func GetPwdById(UserId int64) (pwd string, err error) {
	row := DB.QueryRow(" select password from user_basics where id=?", UserId)
	if err = row.Err(); row.Err() != nil {
		return
	}
	u0 := new(model.UserBasic)
	err = row.Scan(&u0.Password)
	return u0.Password, nil
}
