package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// Md5Encode 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// MD5Encode 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// MakePassword 加密
func MakePassword(plainPassword, salt string) string {
	return Md5Encode(plainPassword + salt)
}

// ValidPassword 解密
func ValidPassword(plainPassword, salt string, password string) bool {
	md := Md5Encode(plainPassword + salt)
	fmt.Println(md + "        " + password)
	return md == password
}
