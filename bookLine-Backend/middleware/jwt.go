package middleware

import (
	"bookLine-Backend/conf"
	"bookLine-Backend/serializer"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("AuthToken")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status":  999,
				"success": false,
				"error":   "请求token为空",
			})
			c.Abort()
			return
		}
		claims, err := JWTParse(token)
		if err != nil || claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusOK, gin.H{
				"status":  999,
				"success": false,
				"error":   "token验证失败,token可能被篡改或token已过期",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func JWTParse(token string) (*serializer.JWTClaims, error) { //解析JWT
	//TODO:看不懂这句话
	tmp, err := jwt.ParseWithClaims(token, &serializer.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Config.Admin.Secret), nil
	})
	if tmp != nil {
		claims, result := tmp.Claims.(*serializer.JWTClaims)
		if result && tmp.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func JWTGenerate(claims serializer.JWTClaims) (string, error) { //生成JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Config.Admin.Secret))
}
