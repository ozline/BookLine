package api

import (
	"bookLine-Backend/conf"
	"bookLine-Backend/middleware"
	"bookLine-Backend/serializer"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func checkAdmin(c *gin.Context) bool { //当部分权限需要管理员来操作时
	token := c.Request.Header.Get("AuthToken")
	claims, _ := middleware.JWTParse(token)
	if claims.IsAdmin != "1" {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"success": false,
			"error":   "权限错误",
		})
		return false
	}
	return true
}

func normalReturn(c *gin.Context, data interface{}, err error, success bool, token string) {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"success": false,
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"success": success,
			"data":    data,
			"token":   token,
		})
	}
}

func errReturn(c *gin.Context, err string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  500,
		"success": false,
		"error":   err,
	})
}

func getAuthToken(id string, username string, isadmin string) (string, error) {
	claims := serializer.JWTClaims{
		Id:       id,
		Username: username,
		IsAdmin:  isadmin,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 3600), //签名生效时间 一小时前
			ExpiresAt: int64(time.Now().Unix() + 3600), //签名过期时间 按一小时算吧
			Issuer:    conf.Config.Admin.Name,
		},
	}
	token, err := middleware.JWTGenerate(claims)
	return token, err
}

func updateAuthToken(c *gin.Context) string {
	before := c.Request.Header.Get("AuthToken")
	if before == "" {
		return ""
	}
	claims, err := middleware.JWTParse(before)
	if err != nil {
		return ""
	}
	token, err := getAuthToken(claims.Id, claims.Username, claims.IsAdmin)
	if err != nil {
		return ""
	} else {
		return token
	}
}
