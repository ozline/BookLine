package model

import (
	"bookLine-Backend/conf"
	"bookLine-Backend/middleware"
	"fmt"
	"os"
	"path"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
)

var Bucket *oss.Bucket

func OSSInit() {
	fmt.Println("OSS SDK VERSION:", oss.Version)
	Client, err := oss.New(conf.Config.OSS.Endpoint, conf.Config.OSS.AccessKeyID, conf.Config.OSS.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		fmt.Println("Error: OSSInit", err)
		os.Exit(-1)
	}
	Bucket, err = Client.Bucket(conf.Config.OSS.Bucket)
	if err != nil {
		fmt.Println("Error: OSSInit", err)
	}
	fmt.Println("Aliyun OSS连接成功")
}

func OSSUpload(objectKey string, filePath string) bool { //上传文件这块可以优化一下
	err := Bucket.PutObjectFromFile(conf.Config.OSS.MainDir+"/"+objectKey, filePath)
	if err != nil {
		fmt.Println("Error: OSSUplaod ", err)
		return false
	}
	return true
}

//RETURN: 上传是否成功，fileName
func OSSUPloadFormFile(c *gin.Context, formname string, dir string) (bool, string) {
	file, err := c.FormFile(formname)
	if err != nil {
		return false, ""
	}

	suffix := path.Ext(file.Filename)
	prefixBefore := middleware.GenerateMD5Random()
	err = c.SaveUploadedFile(file, prefixBefore+suffix)
	if err != nil {
		return false, ""
	}
	prefixAfter := middleware.GenerateMD5(prefixBefore) //TODO:这里对文件进行MD5的话单一性更强
	result := OSSUpload(dir+"/"+prefixAfter+suffix, prefixBefore+suffix)
	os.Remove(prefixBefore + suffix) //删除本地文件
	return result, prefixAfter + suffix
}
