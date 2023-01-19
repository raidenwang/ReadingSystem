package api

import (
	"github.com/gin-gonic/gin"
	"go_code/ReadingSystem/dao"
	"go_code/ReadingSystem/jwt"
	"go_code/ReadingSystem/model"
	"go_code/ReadingSystem/service"
	"net/http"
)

func GetCommentByBookId(c *gin.Context) {
	BookId := c.Param("book_id")
	num := dao.GetCommentNumByBookId(BookId)
	var res = model.SqlResponse{}
	var err error
	var comment [100]model.ApiComment
	comment, err = service.GetCommentByBookId(BookId, dao.User00.ID)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "获取书评错误"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	if num == 0 {
		res.Status = http.StatusBadRequest
		res.Info = "暂无书评" + BookId
		res.Data = "期待您的发言"
		c.JSON(http.StatusOK, res)
		return
	}
	data := comment[0:num]
	res.Status = 10000
	res.Info = "获取书评成功"
	res.Data = data
	c.JSON(http.StatusOK, res)
	return
}
func WriteComment(c *gin.Context) {
	BookId := c.Param("book_id")
	var c0 model.PostComment
	err := c.Bind(&c0)
	if err != nil {
		return
	}
	var res = model.SqlResponse{}
	_, err1 := jwt.ParseToken(jwt.Token)
	if err1 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token失效"
		res.Data = "请重新登录"
		c.JSON(http.StatusOK, res)
		return
	}
	_, err2, data := service.WriteComment(BookId, c0.Content, dao.User00.ID)
	if err2 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "发表书评失败"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Status = 10000
	res.Info = "success"
	res.Data = data
	c.JSON(http.StatusOK, res)
	return
}
func DeleteComment(c *gin.Context) {
	CommentID := c.Param("comment_id")
	var c0 model.PostCommentID
	err := c.Bind(&c0)
	if err != nil {
		return
	}
	var res = model.SqlResponse{}
	_, err1 := jwt.ParseToken(jwt.Token)
	if err1 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token失效"
		res.Data = "请重新登录"
		c.JSON(http.StatusOK, res)
		return
	}
	var num1 int64
	_, err, num1 = service.DeleteComment(CommentID)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "删除书评失败"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Status = 10000
	res.Info = "success"
	res.Data = num1
	c.JSON(http.StatusOK, res)
	return
}
func UpdateComment(c *gin.Context) {
	CommentId := c.Param("comment_id")
	var c0 model.PostComment
	err := c.Bind(&c0)
	if err != nil {
		return
	}
	var res = model.SqlResponse{}
	_, err1 := jwt.ParseToken(jwt.Token)
	if err1 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token失效"
		res.Data = "请重新登录"
		c.JSON(http.StatusOK, res)
		return
	}
	err1 = service.UpdateComment(CommentId, c0.Content)
	if err1 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "书评更新失败"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Status = 10000
	res.Info = "success"
	res.Data = "书评更新成功"
	c.JSON(http.StatusOK, res)
	return

}
