package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zaiyun.app/app/middleware"
	"zaiyun.app/app/models"
)

func GetUserOrders(context *gin.Context) {
	pageCodeStr := context.Query("pageCode")
	if pageCodeStr == "" {
		pageCodeStr = "1"
	}
	pageCode, _ := strconv.Atoi(pageCodeStr)
	pageSizeStr := context.Query("pageSize")
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}
	pageSize, _ := strconv.Atoi(pageSizeStr)
	orders, err := models.GetUserOrders(context.Keys["claims"].(*middleware.MyClaims).UserID, pageCode, pageSize)
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success", "data": orders})
}
