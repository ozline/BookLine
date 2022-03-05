package model

import (
	"bookLine-Backend/conf"
	"bookLine-Backend/middleware"
	"bookLine-Backend/serializer"
	"errors"
	"fmt"
	"strconv"
)

func AddNovel(info serializer.Novel) (bool, error) {
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	cmd := "INSERT INTO `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Data + "` (`title`,`author`,`status`,`filepath`,`coverpath`,`profile`,`category`,`updatetime`,`user`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	tmp, err := tx.Prepare(cmd)
	if err != nil {
		return false, err
	}
	_, err = tmp.Exec(info.Title, info.Author, info.Status, info.FilePath, info.CoverPath, info.Profile, info.Category, info.Update, info.User)
	if err == nil {
		tx.Commit()
		return true, nil
	} else {
		return false, err
	}
}

func GetNovelAll(page string, mode string) (bool, serializer.NovelReturn, error) {
	var array = make([]serializer.Novel, 0)
	var result serializer.NovelReturn
	limitPage, _ := strconv.Atoi(page)
	limitStart := strconv.Itoa((limitPage - 1) * 10)

	cmd := "SELECT * from `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Data + "`"
	cmd += " WHERE `status`=" + mode
	cmd += " ORDER BY id desc LIMIT 10 OFFSET " + limitStart

	fmt.Println(cmd)
	rows, err := db.Query(cmd)
	if err != nil {
		return false, result, err
	}
	for rows.Next() {
		var info serializer.Novel
		err := rows.Scan(&info.Id, &info.Title, &info.Author, &info.Status, &info.FilePath, &info.CoverPath, &info.Profile, &info.Category, &info.Update, &info.User)
		if err != nil {
			return false, result, err
		}
		array = append(array, info)
	}
	cmd = "SELECT count(*) from `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Data + "`"
	cmd += " WHERE `status`=" + mode
	var count = 0
	err = db.QueryRow(cmd).Scan(&count)
	if err != nil {
		return false, result, err
	} else {
		result.Items = array
		result.Total = count
		return true, result, nil
	}
}

func GetNovelCategory(page string, catrgory string) (bool, serializer.NovelReturn, error) {
	var array = make([]serializer.Novel, 0)
	var result serializer.NovelReturn
	limitPage, _ := strconv.Atoi(page)
	limitStart := strconv.Itoa((limitPage - 1) * 10)

	cmd := "SELECT * from `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Data + "`"
	cmd += " WHERE `status`=1 AND `category`=" + catrgory
	cmd += " ORDER BY id desc LIMIT 10 OFFSET " + limitStart

	rows, err := db.Query(cmd)
	if err != nil {
		return false, result, err
	}
	for rows.Next() {
		var info serializer.Novel
		err := rows.Scan(&info.Id, &info.Title, &info.Author, &info.Status, &info.FilePath, &info.CoverPath, &info.Profile, &info.Category, &info.Update, &info.User)
		if err != nil {
			return false, result, err
		}
		array = append(array, info)
	}
	cmd = "SELECT count(*) from `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Data + "`"
	cmd += " WHERE `status`=1 AND `category`=" + catrgory
	var count = 0
	err = db.QueryRow(cmd).Scan(&count)
	if err != nil {
		return false, result, err
	} else {
		result.Items = array
		result.Total = count
		return true, result, nil
	}
}

func GetNovelSingle(id string) (bool, serializer.Novel, error) {
	var info serializer.Novel
	cmd := "SELECT * FROM `" + conf.Config.Mysql.Table.Data + "`"
	cmd += " WHERE `id`=" + id
	err := db.QueryRow(cmd).Scan(&info.Id, &info.Title, &info.Author, &info.Status, &info.FilePath, &info.CoverPath, &info.Profile, &info.Category, &info.Update, &info.User)
	if err != nil {
		return false, info, errors.New("找不到收藏")
	}
	return true, info, nil
}

func SearchNovel(page string, keywords string) (bool, serializer.NovelReturn, error) {
	var array = make([]serializer.Novel, 0)
	var result serializer.NovelReturn
	limitPage, _ := strconv.Atoi(page)
	limitStart := strconv.Itoa((limitPage - 1) * 10)

	cmd := "SELECT * from `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Data + "`"
	cmd += " WHERE `status`='1' AND CONCAT(IFNULL(`title`,''),IFNULL(`author`,''),IFNULL(`profile`,''))"
	cmd += " LIKE '%" + keywords + "%'"
	cmd += " ORDER BY id desc LIMIT 10 OFFSET " + limitStart

	rows, err := db.Query(cmd)
	if err != nil {
		return false, result, err
	}
	for rows.Next() {
		var info serializer.Novel
		err := rows.Scan(&info.Id, &info.Title, &info.Author, &info.Status, &info.FilePath, &info.CoverPath, &info.Profile, &info.Category, &info.Update, &info.User)
		if err != nil {
			return false, result, err
		}
		array = append(array, info)
	}
	cmd = "SELECT count(*) from `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Data + "`"
	cmd += " WHERE `status`='1' AND CONCAT(IFNULL(`title`,''),IFNULL(`author`,''),IFNULL(`profile`,''))"
	cmd += " LIKE '%" + keywords + "%'"
	//TODO: 尽量实现一次查询
	var count = 0
	err = db.QueryRow(cmd).Scan(&count)
	if err != nil {
		return false, result, err
	} else {
		result.Items = array
		result.Total = count
		return true, result, nil
	}
}

func DeleteNovel(id string) (bool, error) {
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	cmd := "DELETE FROM `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Data + "` WHERE `id`=" + id
	tmp, err := tx.Prepare(cmd)
	if err != nil {
		return false, err
	}
	_, err = tmp.Exec()
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}

func UpdateNovel(id string, status string) (bool, error) {
	var cmd string
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	cmd = "UPDATE `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Data + "` SET `status`=" + status + " WHERE `id`=" + id
	tmp, err := tx.Prepare(cmd)
	if err != nil {
		return false, err
	}
	_, err = tmp.Exec()
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}

func FavNovel(userid string, dataid string) (bool, error) {
	res, _ := CheckFav(userid, dataid)
	if res {
		return false, errors.New("已经存在")
	}
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	cmd := "INSERT INTO `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Favlist + "` (`userid`,`novelid`,`time`) VALUES (?, ?, ?)"
	tmp, err := tx.Prepare(cmd)
	if err != nil {
		return false, err
	}
	_, err = tmp.Exec(userid, dataid, strconv.FormatInt(middleware.GetTimestamp(), 10))
	if err == nil {
		tx.Commit()
		return true, nil
	} else {
		return false, err
	}
}

func DeleteFav(userid string, dataid string) (bool, error) {
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	cmd := "DELETE FROM `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Favlist + "` WHERE `userid`=" + userid + " AND `novelid`=" + dataid
	tmp, err := tx.Prepare(cmd)
	if err != nil {
		return false, err
	}
	_, err = tmp.Exec()
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}

func CheckFav(userid string, dataid string) (bool, error) {
	var favid string
	cmd := "SELECT id FROM `" + conf.Config.Mysql.Table.Favlist + "`"
	cmd += " WHERE `userid`=" + userid + " AND `novelid`=" + dataid
	err := db.QueryRow(cmd).Scan(&favid)
	if err != nil {
		return false, errors.New("找不到收藏")
	}
	return true, nil
}

func GetFav(userid string, page string) (bool, serializer.NovelReturn, error) {
	var array = make([]serializer.NovelFav, 0)
	var result serializer.NovelReturn
	limitPage, _ := strconv.Atoi(page)
	limitStart := strconv.Itoa((limitPage - 1) * 10)

	cmd := "SELECT `novelid`,`time` from `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Favlist + "`"
	cmd += " WHERE `userid`=" + userid
	cmd += " ORDER BY id desc LIMIT 10 OFFSET " + limitStart

	rows, err := db.Query(cmd)
	if err != nil {
		return false, result, err
	}
	for rows.Next() {
		var novelid string
		var tmp serializer.NovelFav
		var favtime string
		err := rows.Scan(&novelid, &favtime)
		if err != nil {
			return false, result, err
		}
		_, novel, err := GetNovelSingle(novelid)
		if err != nil {
			return false, result, err
		}
		tmp.Item = novel
		tmp.Favtime = favtime
		array = append(array, tmp)
	}
	cmd = "SELECT count(*) from `" + conf.Config.Mysql.DBName + "`.`" + conf.Config.Mysql.Table.Favlist + "`"
	cmd += " WHERE `userid`=" + userid
	var count = 0
	err = db.QueryRow(cmd).Scan(&count)
	if err != nil {
		return false, result, err
	} else {
		result.Items = array
		result.Total = count
		return true, result, nil
	}
}
