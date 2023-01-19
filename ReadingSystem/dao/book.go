package dao

import (
	"go_code/ReadingSystem/model"
	"log"
	"strings"
)

func GetBookNum() int64 {
	var num int64
	for i := 1; ; i++ {
		b0 := new(model.BookBasic)
		row := DB.QueryRow("select id from book_basics where id=?", i)
		if err := row.Err(); row.Err() != nil {
			log.Print(err)
			break
		}
		err := row.Scan(&b0.Id)
		if err != nil {
			break
		}
		num++
	}
	return num
}

func GetBookUnLogin(num int64) (data [10]model.BookBasic, err error) {
	var i int64
	for i = 0; i < num; i++ {
		row := DB.QueryRow("select * from book_basics where id=?", i+1)
		err = row.Scan(&data[i].Id, &data[i].Name, &data[i].IsStar, &data[i].Author, &data[i].CommentNum, &data[i].Score,
			&data[i].Cover, &data[i].PublishTime, &data[i].Link, &data[i].Label)
		if err != nil {
			log.Println(err)
			return
		}
	}
	return
}

func GetBookLogin(num int64, UserId int64) (data [10]model.BookBasic, err error) {
	var i int64
	for i = 0; i < num; i++ {
		row := DB.QueryRow("select * from book_basics where id=?", i+1)
		err = row.Scan(&data[i].Id, &data[i].Name, &data[i].IsStar, &data[i].Author, &data[i].CommentNum, &data[i].Score,
			&data[i].Cover, &data[i].PublishTime, &data[i].Link, &data[i].Label)
		if err != nil {
			log.Println(err)
			return
		}
		judge := JudgeIsStar(UserId, data[i].Id)
		if judge == UserId {
			data[i].IsStar = true
		}
	}
	return
}

func JudgeIsStar(UserId int64, BookId int64) int64 {
	rows, err := DB.Query("select * from favour_lists where book_id= ?", BookId)
	if err != nil {
		log.Println(err)
		return 0
	}
	var f model.FavourList
	for rows.Next() {
		err = rows.Scan(&f.Id, &f.BookId, &f.UserId)
		if err != nil {
			break
		}
		if f.UserId == UserId {
			return f.UserId
		}
	}
	return 0

}

func SearchBookByName(BookName string, UserId int64) (u model.BookBasic, err error) {
	row := DB.QueryRow("select * from book_basics where name=?", BookName)
	err = row.Scan(&u.Id, &u.Name, &u.IsStar, &u.Author, &u.CommentNum, &u.Score,
		&u.Cover, &u.PublishTime, &u.Link, &u.Label)
	u.IsStar = false
	if err != nil {
		log.Println(err)
		return
	}
	if UserId == 0 {
		return
	}
	judge := JudgeIsStar(UserId, u.Id)
	if judge == UserId {
		u.IsStar = true
	}
	return
}

func StarBook(BookId int, UserId int64) (u model.FavourList, result int) {
	row := DB.QueryRow("select id,book_id,user_id from favour_lists where book_id=? and user_id=?", BookId, UserId)
	if err := row.Err(); row.Err() != nil {
		log.Println(err)
		return
	}
	err := row.Scan(&u.Id, &u.BookId, &u.UserId)
	if u.UserId == UserId {
		result = 1
		return
	}
	sqlStr := "insert into favour_lists(book_id,user_id) values (?,?)"
	_, err = DB.Exec(sqlStr, BookId, UserId)
	if err != nil {
		result = 2
		return
	}
	result = 3
	return
}

func GetBookByLabel(Label string, UserId int64, num int) (data [10]model.BookBasic, err error, n int) {
	for i := 0; i < num; i++ {
		row := DB.QueryRow("select * from book_basics where id=?", i+1)
		var b0 model.BookBasic
		err = row.Scan(&b0.Id, &b0.Name, &b0.IsStar, &b0.Author, &b0.CommentNum, &b0.Score,
			&b0.Cover, &b0.PublishTime, &b0.Link, &b0.Label)
		if err != nil {
			log.Println(err)
			return
		}
		str1 := []rune(b0.Label)
		str2 := []rune(Label)
		if strings.Contains(string(str1), string(str2)) {
			judge := JudgeIsStar(UserId, b0.Id)
			if judge == UserId {
				b0.IsStar = true
			}
			data[n] = b0
			n++
		}
	}
	return
}
