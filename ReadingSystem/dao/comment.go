package dao

import (
	"database/sql"
	"go_code/ReadingSystem/model"
	"log"
	"strconv"
	"time"
)

func GetCommentByBookId(BookId string, UserId int64) (comment [100]model.ApiComment, err error) {

	rows, err := DB.Query("select * from comments where book_id= ?", BookId)
	if err != nil {
		log.Println(err)
		return
	}
	var i int
	for rows.Next() {
		var c model.ApiComment
		var c1 model.Comment
		err = rows.Scan(&c1.Id, &c1.BookId, &c1.PublishTime, &c1.Content, &c1.UserId, &c1.Avatar, &c1.NickName,
			&c1.PraiseCount, &c1.IsPraised, &c1.IsFocus)
		c.BookId = c1.BookId
		c.PublishTime = c1.PublishTime
		c.Content = c1.Content
		c.UserId = c1.UserId
		c.Avatar = c1.Avatar
		c.NickName = c1.NickName
		c.PraiseCount = c1.PraiseCount
		c.IsPraised = c1.IsPraised
		c.IsFocus = c1.IsFocus
		if err != nil {
			log.Println(err)
			return
		}
		if JudgeIsPraised(c1.Id, UserId) {
			c.IsPraised = true
		}
		if JudgeIsFocused(UserId, c1.UserId) {
			c.IsFocus = true
		}
		comment[i] = c
		i++
	}
	return
}

func WriteComment(BookId string, Content string, UserId int64) (result sql.Result, err error, id int64) {
	u, _ := SearchUserById(UserId)
	now := time.Now()
	t := ""
	t += strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "-" + strconv.Itoa(now.Day())
	err = AddCommentNum(BookId)
	if err != nil {
		return
	}
	result, err = DB.Exec("INSERT into comments(book_id,publish_time,content,user_id,avatar,nick_name,praise_count)values (?,?,?,?,?,?,?)",
		BookId, t, Content, u.ID, u.Avatar, u.NickName, 0)
	if err != nil {
		log.Println(err)
		return
	}
	id, _ = result.LastInsertId()
	return
}

func AddCommentNum(BookId string) error {
	id, _ := strconv.Atoi(BookId)
	row := DB.QueryRow("select comment_num from book_basics where id=?", id)
	var b model.BookBasic
	err := row.Scan(&b.CommentNum)
	if err != nil {
		log.Println(err)
		return err
	}
	num := b.CommentNum + 1
	_, err1 := DB.Exec("update book_basics set comment_num=? where id=?", num, id)
	if err != nil {
		log.Println(err1)
		return err1
	}
	return nil
}

func SubCommentNum(BookId string, num int64) error {
	id, _ := strconv.Atoi(BookId)
	row := DB.QueryRow("select comment_num from book_basics where id=?", id)
	var b model.BookBasic
	err := row.Scan(&b.CommentNum)
	if err != nil {
		log.Println(err)
		return err
	}
	newnum := b.CommentNum - num
	_, err1 := DB.Exec("update book_basics set comment_num=? where id=?", newnum, id)
	if err != nil {
		log.Println(err1)
		return err1
	}
	return nil
}

func DeleteComment(CommentID string) (result sql.Result, err error, num1 int64) {
	row := DB.QueryRow("select book_id from comments where id=?", CommentID)
	var c model.Comment
	err = row.Scan(&c.BookId)
	if err != nil {
		log.Println(err)
		return
	}
	result, err = DB.Exec("delete from comments where id=?", CommentID)
	if err != nil {
		log.Println(err)
		return
	}
	num, _ := result.RowsAffected()
	num1 = num
	err = SubCommentNum(c.BookId, num)
	if err != nil {
		return
	}
	err = UpdateIsPraised(CommentID)
	if err != nil {
		return
	}
	return
}

func GetCommentNumByBookId(BookId string) int64 {
	var num int64
	rows, err := DB.Query("select id from comments where book_id= ?", BookId)
	if err != nil {
		log.Println(err)
		return 0
	}
	var c0 model.Comment
	for rows.Next() {
		err := rows.Scan(&c0.Id)
		if err != nil {
			log.Println(err)
			return num
		}
		num++
	}
	return num
}

func JudgeIsPraised(CommentId int64, UserId int64) bool {
	rows, err := DB.Query("select * from praised_lists where comment_id= ?", CommentId)
	if err != nil {
		log.Println(err)
		return false
	}
	var f model.PraisedList
	for rows.Next() {
		err = rows.Scan(&f.Id, &f.CommentId, &f.UserId)
		if err != nil {
			break
		}
		if f.UserId == UserId {
			return true
		}
	}
	return false

}

func JudgeIsFocused(UserId int64, FocusedUserId int64) bool {
	rows, err := DB.Query("select * from focused_lists where user_id= ?", UserId)
	if err != nil {
		log.Println(err)
		return false
	}
	var f model.FocusedList
	for rows.Next() {
		err = rows.Scan(&f.Id, &f.UserId, &f.FocusedUserId)
		if err != nil {
			break
		}
		if f.FocusedUserId == FocusedUserId {
			return true
		}
	}
	return false

}

func UpdateIsPraised(CommentId string) error {
	comment_id, _ := strconv.Atoi(CommentId)
	_, err := DB.Exec("delete from praised_lists where comment_id=?", comment_id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateComment(CommentId string, Content string) error {
	id, _ := strconv.Atoi(CommentId)
	_, err := DB.Exec("update comments set content=? where id=?", Content, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
