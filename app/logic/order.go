package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zaiyun.app/app/models"
)

func GetUserOrders(context *gin.Context) {
	models.GetUserOrders(1)
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "getorders ok"})
}
