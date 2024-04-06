package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrders(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "getorders ok"})
}
