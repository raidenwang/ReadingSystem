package model

type UserBasic struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	NickName     string `json:"nickName"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Telephone    string `json:"telephone"`
	QQ           string `json:"QQ"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Birthday     string `json:"birthday"`
}
type UserInfoShow struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	NickName     string `json:"nickName"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Telephone    string `json:"telephone"`
	QQ           string `json:"QQ"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Birthday     string `json:"birthday"`
}
type ApiUser struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	OldPassword string `json:"oldPassword"`
}

type Pwd struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type SqlResponse struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}
