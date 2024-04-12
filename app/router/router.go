package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"zaiyun.app/app/logic"
	"zaiyun.app/app/middleware"
	"zaiyun.app/app/tools"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/", logic.Index)

	goods := r.Group("/")
	{
		goods.GET("/getGoods", logic.GetGoods)
		goods.GET("/getGoodsDetail", logic.GetGoodsDetail)
	}

	r.GET("/getIndexSwipers", logic.GetIndexSwipers)

	user := r.Group("/")
	{
		user.POST("/login", logic.PostLogin)
		user.POST("/register", logic.CreateUser)
	}

	orders := r.Group("/")
	orders.Use(middleware.Verify)
	{
		orders.GET("/getOrders", logic.GetOrders)
	}

	{
		r.GET("/captcha", func(context *gin.Context) {
			captcha, err := tools.CaptchaGenerate()
			if err != nil {
				context.JSON(http.StatusOK, tools.ECode{
					Code:    10005,
					Message: err.Error(),
				})
				return
			}

			context.JSON(http.StatusOK, tools.ECode{
				Data: captcha,
			})
		})

		r.POST("/captcha/verify", func(context *gin.Context) {
			var param tools.CaptchaData
			if err := context.ShouldBind(&param); err != nil {
				context.JSON(http.StatusOK, tools.ParamErr)
				return
			}

			fmt.Printf("参数为：%+v", param)
			if !tools.CaptchaVerify(param) {
				context.JSON(http.StatusOK, tools.ECode{
					Code:    10008,
					Message: "验证失败",
				})
				return
			}
			context.JSON(http.StatusOK, tools.OK)
		})
	}

	if err := r.Run(":8090"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
