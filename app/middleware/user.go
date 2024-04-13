package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 写成这种形式，方便传递参数
func VerifyMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := context.Cookie("token")
		if err != nil || token == "" {
			context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "未登录"})
			context.Abort()
			return
		}
		context.Next()
	}
}

// 写成这种形式，无法传递参数
func VerifyAdmin(context *gin.Context) {
	token, err := context.Cookie("token")
	if err != nil || token != "admin" {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "未登录"})
		context.Abort()
		return
	}
	context.Next()
}
