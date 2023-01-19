package main

import (
	"go_code/ReadingSystem/api"
	"go_code/ReadingSystem/dao"
)

func main() {
	dao.InitDB()

	api.Router()
}
