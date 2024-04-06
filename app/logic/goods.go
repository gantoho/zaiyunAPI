package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zaiyun.app/app/model"
)

func GetGoods(context *gin.Context) {
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
	goods := model.GetGoods(pageCode, pageSize)
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": goods, "total": len(goods)})
}

func GetGoodsDetail(context *gin.Context) {
	idStr := context.Query("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	goods := model.GetGoodsDetail(id)
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": goods})
}
