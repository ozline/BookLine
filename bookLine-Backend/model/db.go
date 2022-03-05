package model

import (
	"bookLine-Backend/conf"
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DBInit() bool { //连接RDS
	path := strings.Join([]string{conf.Config.Mysql.Username, ":", conf.Config.Mysql.Password, "@tcp(", conf.Config.Mysql.Address, ":", conf.Config.Mysql.Port, ")/", conf.Config.Mysql.DBName, "?charset=utf8&parseTime=True"}, "")
	db, _ = sql.Open("mysql", path)
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return true
}

func DBDestory() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
