package api

import (
	"github.com/gin-gonic/gin"
	"go_code/ReadingSystem/dao"
	"go_code/ReadingSystem/jwt"
	"go_code/ReadingSystem/model"
	"go_code/ReadingSystem/service"
	"net/http"
	"strconv"
)

func GetFavourList(c *gin.Context) {
	var res = model.SqlResponse{}
	_, err1 := jwt.ParseToken(jwt.Token)
	if err1 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token失效"
		res.Data = "请重新登录"
		c.JSON(http.StatusOK, res)
		return
	}
	collections, err := service.GetFavourList(dao.User00.ID)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Info = "获取用户收藏列表失败"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Status = 10000
	res.Info = "success"
	res.Data = collections
	c.JSON(http.StatusOK, res)
	return
}

func Focus(c *gin.Context) {
	var res = model.SqlResponse{}
	_, err1 := jwt.ParseToken(jwt.Token)
	if err1 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token失效"
		res.Data = "请重新登录"
		c.JSON(http.StatusOK, res)
		return
	}
	FocusUserId := c.PostForm("user_id")
	FId, _ := strconv.Atoi(FocusUserId)
	result := service.AddFocused(int64(FId), dao.User00.ID)
	if result == 0 {
		res.Status = http.StatusBadRequest
		res.Info = "关注用户失败"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	if result == 1 {
		res.Status = http.StatusBadRequest
		res.Info = "已关注该用户"
		res.Data = "不要重复操作"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Status = 10000
	res.Info = "success·"
	res.Data = "关注成功"
	c.JSON(http.StatusOK, res)
	return
}

func Praise(c *gin.Context) {
	modelnum := c.DefaultPostForm("modelnum", "0")
	targetId := c.DefaultPostForm("target_id", "0")
	var res = model.SqlResponse{}
	_, err1 := jwt.ParseToken(jwt.Token)
	if err1 != nil {
		res.Status = http.StatusBadRequest
		res.Info = "token失效"
		res.Data = "请重新登录"
		c.JSON(http.StatusOK, res)
		return
	}
	if modelnum == "0" || targetId == "0" {
		res.Status = http.StatusBadRequest
		res.Info = "参数获取错误"
		res.Data = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	if modelnum == "1" {
		result := service.PraisePost(targetId, dao.User00.ID)
		if result == 0 {
			res.Status = http.StatusBadRequest
			res.Info = "您已点过赞"
			res.Data = "请勿重复操作"
			c.JSON(http.StatusOK, res)
			return
		}
		if result == 2 {
			res.Status = http.StatusBadRequest
			res.Info = "点赞失败"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		res.Status = 10000
		res.Info = "点赞成功"
		res.Data = targetId
		c.JSON(http.StatusOK, res)
		return
	}
	if modelnum == "2" {
		result := service.PraiseComment(targetId, dao.User00.ID)
		if result == 0 {
			res.Status = http.StatusBadRequest
			res.Info = "您已点过赞"
			res.Data = "请勿重复操作"
			c.JSON(http.StatusOK, res)
			return
		}
		if result == 2 {
			res.Status = http.StatusBadRequest
			res.Info = "点赞失败"
			res.Data = "error"
			c.JSON(http.StatusOK, res)
			return
		}
		res.Status = 10000
		res.Info = "点赞成功"
		res.Data = targetId
		c.JSON(http.StatusOK, res)
		return
	}
}
