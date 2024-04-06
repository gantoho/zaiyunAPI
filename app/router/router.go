package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zaiyun.app/app/logic"
	"zaiyun.app/app/middleware"
)

func InitRouter() {
	r := gin.Default()

	goods := r.Group("/")
	{
		goods.GET("/getGoods", logic.GetGoods)
		goods.GET("/getGoodsDetail", logic.GetGoodsDetail)
	}

	r.GET("/getIndexSwipers", logic.GetIndexSwipers)

	user := r.Group("/")
	{
		user.POST("/login", logic.PostLogin)
	}

	orders := r.Group("/")
	orders.Use(middleware.Verify)
	{
		orders.GET("/getOrders", logic.GetOrders)
	}

	if err := r.Run(":8090"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
