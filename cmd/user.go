package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	"message-board/api"
	"message-board/dao"
)

func main() {
	dao.InitDB()
	api.InitEngine()
}
