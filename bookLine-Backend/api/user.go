package api

import (
	"bookLine-Backend/conf"
	"bookLine-Backend/middleware"
	"bookLine-Backend/model"
	"bookLine-Backend/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) { //用户注册
	var info serializer.User
	info.Username = c.PostForm("username")
	info.Password = c.PostForm("password")

	result, err := model.AddUser(info)
	normalReturn(c, nil, err, result, "")
}

func UserLogin(c *gin.Context) { //用户登录
	var info serializer.User
	info.Username = c.PostForm("username")
	info.Password = c.PostForm("password")

	if len(info.Username) == 0 || len(info.Password) == 0 {
		errReturn(c, "用户名或密码错误")
		return
	}
	tmp, id, isadmin, err := model.CheckUser(info)
	if err != nil {
		errReturn(c, err.Error())
	} else {
		token, err := getAuthToken(id, info.Username, isadmin)
		normalReturn(c, nil, err, tmp, token)
	}
}

func UserTest(c *gin.Context) { //测试通信
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": true,
	})
}

//TODO:可能存在漏洞
func UserUpload(c *gin.Context) { //上传文件
	var upload serializer.Novel
	dirFile := "novel"
	dirCover := "cover"

	//上传文件
	result, fileName := model.OSSUPloadFormFile(c, "file", dirFile)
	if !result {
		errReturn(c, "小说文件传输失败")
		return
	}
	upload.FilePath = "https://" + conf.Config.OSS.Endpoint + "/" + conf.Config.OSS.MainDir + "/" + dirFile + "/" + fileName

	//上传封面
	result, fileName = model.OSSUPloadFormFile(c, "cover", dirCover)
	if !result {
		errReturn(c, "封面文件传输失败")
		return
	}
	upload.CoverPath = "https://" + conf.Config.OSS.Endpoint + "/" + conf.Config.OSS.MainDir + "/" + dirCover + "/" + fileName

	//其他处理
	upload.Title = c.PostForm("title")
	upload.Category = c.PostForm("category")
	upload.Profile = c.PostForm("profile")
	upload.Status = 0
	upload.Author = c.PostForm("author")
	upload.User = c.PostForm("user")
	upload.Update = c.PostForm("time")

	res, err := model.AddNovel(upload)

	path := map[string]string{
		"filePath":  upload.FilePath,
		"coverPath": upload.CoverPath,
	}
	normalReturn(c, path, err, res, updateAuthToken(c))
}

func CheckToken(c *gin.Context) { //核查token
	token := c.Request.Header.Get("AuthToken")
	claims, _ := middleware.JWTParse(token)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": true,
		"data":    claims,
	})
}

func getUserID(c *gin.Context) string {
	token := c.Request.Header.Get("AuthToken")
	claims, _ := middleware.JWTParse(token)
	return claims.Id
}
