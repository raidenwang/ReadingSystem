package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_code/ReadingSystem/model"
	"strings"
)

var (
	DB *sql.DB
)
var User00 = model.ApiUser{
	ID:          0,
	Name:        "",
	OldPassword: "",
}

const (
	name       = "root"
	dbpassword = "123456"
	ip         = "192.168.0.111"
	port       = "3306"
	dbName     = "db1"
)

func InitDB() {
	dsn := strings.Join([]string{name, ":", dbpassword, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//"root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = sql.Open("mysql", dsn)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("连接数据库失败")
		return
	}
	fmt.Println("连接数据库成功")
	return
}
