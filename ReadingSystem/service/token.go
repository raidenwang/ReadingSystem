package service

import (
	"go_code/ReadingSystem/dao"
)

func CreateTokens(Token string, RefreshToken string) error {
	_, err := dao.InsertTokens(Token, RefreshToken)
	return err
}
func SearchToken(RefreshToken string) error {
	_, err := dao.SearchTokensByRefreshToken(RefreshToken)
	return err
}
