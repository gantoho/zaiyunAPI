package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zaiyun.app/app/model"
)

func GetUserOrders(context *gin.Context) {
	model.GetUserOrders(1)
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "getorders ok"})
}
