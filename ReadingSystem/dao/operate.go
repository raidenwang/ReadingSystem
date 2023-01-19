package dao

import (
	"go_code/ReadingSystem/model"
	"log"
	"strconv"
)

func GetFavourList(UserId int64) (data []model.Collection, err error) {
	var data1 []model.Collection
	rows, err1 := DB.Query("select book_id from favour_lists where user_id= ?", UserId)
	if err1 != nil {
		log.Println(err1)
		err = err1
		return
	}
	for rows.Next() {
		var c0 model.Collection
		var c1 model.FavourList
		err = rows.Scan(&c1.BookId)
		if err != nil {
			log.Println(err)
			return
		}
		c0.Id = c1.BookId
		row := DB.QueryRow("select name,publish_time,link from book_basics where id=?", c0.Id)
		err = row.Scan(&c0.Name, &c0.PublishTime, &c0.Link)
		if err != nil {
			log.Println(err)
			return
		}
		data1 = append(data1, c0)
	}
	data = data1
	return
}
func AddFocused(FocusUserId int64, UserId int64) (result int) {
	rows, err := DB.Query("select focused_user_id from focused_lists where user_id= ?", UserId)
	if err != nil {
		log.Println(err)
		return
	}
	var f model.FocusedList
	for rows.Next() {
		err := rows.Scan(&f.FocusedUserId)
		if err != nil {
			log.Println(err)
			return
		}
		if f.FocusedUserId == FocusUserId {
			result = 1
			return
		}
	}
	_, err = DB.Exec("insert into focused_lists (user_id,focused_user_id) value (?,?)", UserId, FocusUserId)
	if err != nil {
		log.Println(err)
		return
	}
	result = 2
	return
}
func PraiseComment(CommentId string, UserId int64) (result int) {
	id, _ := strconv.Atoi(CommentId)
	if JudgeIsPraised(int64(id), UserId) {
		result = 0
		return
	}
	_, err := DB.Exec("insert into praised_lists (comment_id,user_id) value (?,?)", id, UserId)
	if err != nil {
		log.Println(err)
		result = 2
		return
	}
	result = 1
	return
}
func PraisePost(PostId string, UserId int64) (result int) {
	id, _ := strconv.Atoi(PostId)
	if JudgePostIsPraised(int64(id), UserId) {
		result = 0
		return
	}
	_, err := DB.Exec("insert into post_praised_lists (post_id,user_id) value (?,?)", id, UserId)
	if err != nil {
		log.Println(err)
		result = 2
		return
	}
	result = 1
	return
}
func JudgePostIsPraised(PostId int64, UserId int64) bool {
	rows, err := DB.Query("select * from post_praised_lists where post_id= ?", PostId)
	if err != nil {
		log.Println(err)
		return false
	}
	var f model.PostPraisedList
	for rows.Next() {
		err = rows.Scan(&f.Id, &f.PostId, &f.UserId)
		if err != nil {
			break
		}
		if f.UserId == UserId {
			return true
		}
	}
	return false

}
