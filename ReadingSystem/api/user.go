package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/ReadingSystem/dao"
	"go_code/ReadingSystem/jwt"
	"go_code/ReadingSystem/model"
	"go_code/ReadingSystem/service"
	"net/http"
	"strconv"
)

func UserRegister(c *gin.Context) {
	var u model.UserBasic
	err := c.Bind(&u)
	var res = model.SqlResponse{}
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "参数错误"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	u1, _ := service.SearchUserByUserName(u.Name)
	//fmt.Println(u1.Name)
	if u1.Name != "" {
		res.Status = http.StatusBadRequest
		res.Info = "账户已存在"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	sqlStr := "insert into user_basics(name,password) values (?,?)"
	ret, err := dao.DB.Exec(sqlStr, u.Name, u.Password)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		res.Status = http.StatusBadRequest
		res.Info = "写入失败"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Status = 1000
	res.Info = "写入成功"
	res.Data = "OK"
	c.JSON(http.StatusOK, res)
	fmt.Println(ret.LastInsertId()) //打印结果
}

func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var res = model.SqlResponse{}

	code := service.Login(username, password)
	if code == 0 {
		res.Status = http.StatusBadRequest
		res.Info = "cnm"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	if code == 1 {
		res.Status = http.StatusBadRequest
		res.Info = "用户不存在"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	if code == 2 {
		res.Status = http.StatusBadRequest
		res.Info = "密码错误"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	jwt.Token, _ = jwt.GenerateToken(username, password)
	_, err1 := jwt.ParseToken(jwt.Token)
	if err1 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token失效"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	jwt.RefreshToken, _ = jwt.GenerateToken(username, password)
	_, err2 := jwt.ParseToken(jwt.RefreshToken)
	if err2 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "refreshToken失效"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	err3 := service.CreateTokens(jwt.Token, jwt.RefreshToken)
	if err3 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token存入数据库错误"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	data := make(map[string]map[string]string)
	data["data"] = make(map[string]string, 2)
	data["data"]["refresh_token"] = jwt.RefreshToken
	data["data"]["token"] = jwt.Token
	res.Status = 10000
	res.Info = "success"
	if code == 0 {
		res.Data = "cnm"
	}
	res.Data = data
	c.JSON(http.StatusOK, res)
	return
}

func Refresh(c *gin.Context) {
	var u model.Tokens
	err := c.Bind(&u)
	var res = model.SqlResponse{}
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "参数错误"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	NewToken, _ := jwt.GenerateToken("", "")
	NewRefreshToken, _ := jwt.GenerateToken("", "")
	_, err1 := dao.DB.Exec("update tokens set token=?,refresh_token=? where refresh_token=?", NewToken, NewRefreshToken, jwt.RefreshToken)
	if err1 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token更新失败"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	jwt.Token = NewToken
	jwt.RefreshToken = NewRefreshToken
	data := make(map[string]map[string]string)
	data["data"] = make(map[string]string, 2)
	data["data"]["refresh_token"] = jwt.RefreshToken
	data["data"]["token"] = jwt.Token
	res.Status = 10000
	res.Info = "success"
	res.Data = data
	c.JSON(http.StatusOK, res)
	return
}

func UpdatePassword(c *gin.Context) {

	var res = model.SqlResponse{}
	_, err := jwt.ParseToken(jwt.Token)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token失效"
		res.Data = "请刷新界面"
		c.JSON(http.StatusOK, res)
		return
	}
	oldPassword := c.PostForm("oldPassword")
	var newPassword = c.PostForm("newPassword")
	err = c.Bind(&oldPassword)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "oldPassword参数错误"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	err = c.Bind(&newPassword)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "newPassword参数错误"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	pwd, _ := service.GetPwdById(dao.User00.ID)
	if oldPassword != pwd {
		res.Status = http.StatusBadRequest
		res.Info = "密码错误"
		res.Data = "err"
		c.JSON(http.StatusOK, res)
		return
	}

	ret, err3 := dao.DB.Exec("update user_basics set password=? where id=?", newPassword, dao.User00.ID)
	if err3 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "密码更新失败"
		res.Data = "500"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Status = 10000
	res.Info = "success"
	res.Data = "密码修改成功"
	c.JSON(http.StatusOK, res)
	fmt.Println(ret.LastInsertId())
	return
}

func GetUserDate(c *gin.Context) {
	userid := c.Param("user_id")
	id, _ := strconv.Atoi(userid)
	var res model.SqlResponse
	u, err := service.SearchUserById(int64(id))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "ID获取错误"
		res.Data = id
		c.JSON(http.StatusOK, res)
		return
	}
	var user model.UserInfoShow
	user.ID = u.ID
	user.Name = u.Name
	user.NickName = u.NickName
	user.Avatar = u.Avatar
	user.Introduction = u.Introduction
	user.Gender = u.Gender
	user.QQ = u.QQ
	user.Email = u.Email
	user.Telephone = u.Telephone
	user.Birthday = u.Birthday
	var Info = make(map[string]model.UserInfoShow, 10)
	Info["user"] = user
	res.Status = http.StatusOK
	res.Info = "success"
	res.Data = Info
	c.JSON(http.StatusOK, res)
	return
}

func UpdateUserDate(c *gin.Context) {

	var res = model.SqlResponse{}
	_, err := jwt.ParseToken(jwt.Token)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token失效"
		res.Data = "请刷新界面"
		c.JSON(http.StatusOK, res)
		return
	}
	index := 0
	NewNickName := c.PostForm("nickname")
	if NewNickName != "" {
		_, err := dao.DB.Exec("update user_basics set nick_name=? where id=?", NewNickName, dao.User00.ID)
		if err != nil {
			res.Status = http.StatusBadRequest
			res.Info = "昵称更新失败"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		index++
	}
	NewAvatar := c.PostForm("avatar")
	if NewAvatar != "" {
		_, err := dao.DB.Exec("update user_basics set avatar=? where id=?", NewAvatar, dao.User00.ID)
		if err != nil {
			res.Status = http.StatusBadRequest
			res.Info = "头像更新失败"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		index++
	}
	NewIntroduction := c.PostForm("introduction")
	if NewIntroduction != "" {
		_, err := dao.DB.Exec("update user_basics set introduction=? where id=?", NewIntroduction, dao.User00.ID)
		if err != nil {
			res.Status = http.StatusBadRequest
			res.Info = "简介更新失败"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		index++
	}
	NewTelephone := c.PostForm("telephone")
	if NewTelephone != "" {
		_, err := dao.DB.Exec("update user_basics set telephone=? where id=?", NewTelephone, dao.User00.ID)
		if err != nil {
			res.Status = http.StatusBadRequest
			res.Info = "电话更新失败"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		index++
	}
	NewQQ := c.PostForm("qq")
	if NewQQ != "" {
		_, err := dao.DB.Exec("update user_basics set qq=? where id=?", NewQQ, dao.User00.ID)
		if err != nil {
			res.Status = http.StatusBadRequest
			res.Info = "QQ更新失败"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		index++
	}
	NewGender := c.PostForm("gender")
	if NewGender != "" {
		_, err := dao.DB.Exec("update user_basics set gender=? where id=?", NewGender, dao.User00.ID)
		if err != nil {
			res.Status = http.StatusBadRequest
			res.Info = "性别更新失败"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		index++
	}
	NewEmail := c.PostForm("email")
	if NewEmail != "" {
		_, err := dao.DB.Exec("update user_basics set email=? where id=?", NewEmail, dao.User00.ID)
		if err != nil {
			res.Status = http.StatusBadRequest
			res.Info = "邮箱更新失败"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		index++
	}
	NewBirthday := c.PostForm("birthday")
	if NewBirthday != "" {
		_, err := dao.DB.Exec("update user_basics set birthday=? where id=?", NewBirthday, dao.User00.ID)
		if err != nil {
			res.Status = http.StatusBadRequest
			res.Info = "生日更新失败"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		index++
	}
	if index == 0 {
		res.Status = http.StatusBadRequest
		res.Info = "未更新任何信息"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Status = 10000
	res.Info = "success"
	res.Data = "个人信息更新成功"
	c.JSON(http.StatusOK, res)
	return
}
