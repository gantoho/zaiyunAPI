package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Verify(context *gin.Context) {
	token, err := context.Cookie("token")
	if err != nil || token == "" {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "未登录"})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "已登录"})
	context.Next()
}
