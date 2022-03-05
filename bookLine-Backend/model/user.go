package model

import (
	"bookLine-Backend/conf"
	"bookLine-Backend/middleware"
	"bookLine-Backend/serializer"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func AddUser(info serializer.User) (bool, error) {
	var tmplory string
	cmd := "SELECT password FROM `" + conf.Config.Mysql.Table.Users + "` where username = '" + info.Username + "'"
	err := db.QueryRow(cmd).Scan(&tmplory)
	if err == nil {
		return false, errors.New("用户已经存在")
	}

	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	cmd = "INSERT INTO `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Users + "` (`username`,`password`) VALUES (?, ?)"
	tmp, err := tx.Prepare(cmd)
	if err != nil {
		return false, err
	}
	_, err = tmp.Exec(info.Username, middleware.GenerateTokenSHA256(info.Password))
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}

func CheckUser(info serializer.User) (bool, string, string, error) {
	var password string
	var id string
	var isadmin string
	cmd := "SELECT password,id,isadmin FROM `" + conf.Config.Mysql.Table.Users + "` where username = '" + info.Username + "'"
	err := db.QueryRow(cmd).Scan(&password, &id, &isadmin)
	if err != nil {
		return false, "-1", "0", err
	}
	if middleware.GenerateTokenSHA256(info.Password) == password {
		return true, id, isadmin, nil
	} else {
		return false, "-1", "0", nil
	}
}
