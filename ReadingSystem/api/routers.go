package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Router() {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/register", UserRegister)
		user.GET("/token", UserLogin)
		user.GET("/token/refresh", Refresh)
		user.PUT("/password", UpdatePassword)
		user.GET("/info/:user_id", GetUserDate)
		user.PUT("/info", UpdateUserDate)
	}
	book := r.Group("/book")
	{
		book.GET("/list", GetBookList)
		book.GET("/search", SearchBookByName)
		book.PUT("/star", StarBook)
		book.GET("/label", GetBookByLabel)
	}
	comment := r.Group("/comment")
	{
		comment.GET("/:book_id", GetCommentByBookId)
		comment.POST("/:book_id", WriteComment)
		comment.DELETE("/:comment_id", DeleteComment)
		comment.PUT("/:comment_id", UpdateComment)
	}
	operator := r.Group("/operate")
	{
		operator.GET("/collect/list", GetFavourList)
		operator.PUT("/focus", Focus)
		operator.PUT("/praise", Praise)
	}
	err := r.Run(":9090")
	if err != nil {
		log.Printf("%v", err)
		return
	}

}
