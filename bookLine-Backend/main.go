package main

import (
	"bookLine-Backend/conf"
	"bookLine-Backend/model"
	"bookLine-Backend/routes"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	conf.Init()
	model.OSSInit()
	if model.DBInit() {
		r := routes.NewRouter()
		_ = r.Run("localhost:" + conf.Config.HttpPort)
	} else {
		fmt.Println("启动失败")
	}
}

func main() {
	Init()
}
