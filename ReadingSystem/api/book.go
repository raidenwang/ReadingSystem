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

func GetBookList(c *gin.Context) {
	num := service.GetBookNum()
	var res = model.SqlResponse{}
	var err error
	if dao.User00.ID == 0 {
		var data [10]model.BookBasic
		data, err = service.GetBookUnLogin(num)
		if err != nil {
			res.Status = http.StatusBadRequest
			res.Info = "获取书籍列表错误"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		book := data[0:num]
		res.Status = 10000
		res.Info = "获取书籍列表成功（游客访问)"
		res.Data = book
		c.JSON(http.StatusOK, res)
		return
	} else {
		var data [10]model.BookBasic
		data, err = service.GetBookLogin(num, dao.User00.ID)
		if err != nil {
			res.Status = http.StatusBadRequest
			res.Info = "获取书籍列表错误"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		book := data[0:num]
		res.Status = 10000
		res.Info = "获取书籍列表成功"
		res.Data = book
		c.JSON(http.StatusOK, res)
		fmt.Println(data)
		return
	}
}

func SearchBookByName(c *gin.Context) {
	BookName := c.PostForm("BookName")
	var res = model.SqlResponse{}
	_, err := jwt.ParseToken(jwt.Token)
	if err != nil {
		u, err1 := service.SearchBookByName(BookName, 0)
		if err1 != nil {
			res.Status = http.StatusBadRequest
			res.Info = "未查到该书籍(游客模式)"
			res.Data = "请重新输入书籍名称"
			c.JSON(http.StatusOK, res)
			return
		}
		res.Status = 10000
		res.Info = "success(游客模式)"
		res.Data = u
		c.JSON(http.StatusOK, res)
		return
	} else {
		u, err1 := service.SearchBookByName(BookName, dao.User00.ID)
		if err1 != nil {
			res.Status = http.StatusBadRequest
			res.Info = "未查到该书籍"
			res.Data = "请重新输入书籍名称"
			c.JSON(http.StatusOK, res)
			return
		}
		res.Status = 10000
		res.Info = "success"
		res.Data = u
		c.JSON(http.StatusOK, res)
		return
	}
}

func StarBook(c *gin.Context) {
	bookId := c.PostForm("bookId")
	BookId, _ := strconv.Atoi(bookId)
	var res = model.SqlResponse{}
	_, err := jwt.ParseToken(jwt.Token)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token失效"
		res.Data = "请刷新界面"
		c.JSON(http.StatusOK, res)
		return
	}
	_, result := service.StarBook(BookId, dao.User00.ID)
	if result == 0 {
		res.Status = http.StatusBadRequest
		res.Info = "发生错误"
		res.Data = "请上报错误"
		c.JSON(http.StatusOK, res)
		return
	}
	if result == 1 {
		res.Status = http.StatusBadRequest
		res.Info = "您已经收藏过该书籍，请勿重复操作"
		res.Data = 400
		c.JSON(http.StatusOK, res)
		return
	}
	if result == 2 {
		res.Status = http.StatusBadRequest
		res.Info = "收藏失败"
		res.Data = 500
		c.JSON(http.StatusOK, res)
		return
	}
	res.Status = 10000
	res.Info = "success"
	res.Data = "收藏成功"
	c.JSON(http.StatusOK, res)
	return
}

func GetBookByLabel(c *gin.Context) {
	Label := c.Query("label")
	//Label := c.PostForm("label")
	var res = model.SqlResponse{}
	var book [10]model.BookBasic
	num := service.GetBookNum()
	book, err, n := service.GetBookByLabel(Label, dao.User00.ID, int(num))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "获取书籍列表失败"
		res.Data = "请重新输入标签Label"
		c.JSON(http.StatusOK, res)
		return
	}
	data := book[0:n]
	res.Status = 10000
	res.Info = "success"
	res.Data = data
	c.JSON(http.StatusOK, res)
	return
}
