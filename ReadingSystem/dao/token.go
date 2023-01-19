package dao

import (
	"database/sql"
	"go_code/ReadingSystem/model"
)

func InsertTokens(Token string, RefreshToken string) (result sql.Result, err error) {
	result, err = DB.Exec("INSERT into tokens(token,refresh_token)values (?,?)", Token, RefreshToken)
	return result, err
}

func SearchTokensByRefreshToken(RefreshToken string) (u model.Tokens, err error) {
	row := DB.QueryRow("select *from tokens where refresh_token=?", RefreshToken)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.RefreshToken)
	return
}
