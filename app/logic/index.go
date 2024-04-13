package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zaiyun.app/app/models"
)

func GetIndexSwipers(context *gin.Context) {
	indexSwipers := models.GetIndexSwipers()
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": indexSwipers, "total": len(indexSwipers)})
}

func Index(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "Hello! Welcome to ZaiYunAPI."})
}
