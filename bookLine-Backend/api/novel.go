package api

import (
	"bookLine-Backend/model"

	"github.com/gin-gonic/gin"
)

func GetNovelReview(c *gin.Context) {
	var page = c.Query("page")
	if !checkAdmin(c) {
		return
	}
	success, result, err := model.GetNovelAll(page, "0")
	normalReturn(c, result, err, success, updateAuthToken(c))
}

func GetNovelAll(c *gin.Context) {
	var page = c.Query("page")
	if !checkAdmin(c) {
		return
	}
	success, result, err := model.GetNovelAll(page, "-1")
	normalReturn(c, result, err, success, updateAuthToken(c))
}

func GetNovelPassed(c *gin.Context) {
	var page = c.Query("page")
	success, result, err := model.GetNovelAll(page, "1")
	normalReturn(c, result, err, success, updateAuthToken(c))
}

func GetNovelCategory(c *gin.Context) {
	var page = c.Query("page")
	var category = c.Query("category")
	success, result, err := model.GetNovelCategory(page, category)
	normalReturn(c, result, err, success, updateAuthToken(c))
}

func DeleteNovel(c *gin.Context) {
	var id = c.Param("id")
	if !checkAdmin(c) {
		return
	}
	success, err := model.DeleteNovel(id)
	normalReturn(c, nil, err, success, updateAuthToken(c))
}

func UpdateNovel(c *gin.Context) { //暂只支持修改审核状态
	var id = c.Param("id")
	var status = c.Param("status")
	if !checkAdmin(c) {
		return
	}
	success, err := model.UpdateNovel(id, status)
	normalReturn(c, nil, err, success, updateAuthToken(c))
}

func SearchNovel(c *gin.Context) {
	var keywords = c.Query("keywords")
	var page = c.Query("page")
	success, result, err := model.SearchNovel(page, keywords)
	normalReturn(c, result, err, success, updateAuthToken(c))
}

func FavNovel(c *gin.Context) {
	var id = c.Param("id")
	result, err := model.FavNovel(getUserID(c), id)
	normalReturn(c, nil, err, result, updateAuthToken(c))
}

func DeleteFav(c *gin.Context) {
	var id = c.Param("id")
	result, err := model.DeleteFav(getUserID(c), id)
	normalReturn(c, nil, err, result, updateAuthToken(c))
}

func GetFav(c *gin.Context) {
	var page = c.Query("page")
	success, result, err := model.GetFav(getUserID(c), page)
	normalReturn(c, result, err, success, updateAuthToken(c))
}
