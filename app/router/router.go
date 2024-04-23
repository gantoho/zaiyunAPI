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
	goods.Use(middleware.JWTAuth())
	{
		goods.GET("/getGoods", logic.GetGoods)
		goods.GET("/getGoodsDetail", logic.GetGoodsDetail)
	}

	r.GET("/getIndexSwipers", logic.GetIndexSwipers)

	login := r.Group("/")
	{
		login.GET("/login", logic.GetLogin)
		login.POST("/login", logic.PostLogin)
		login.POST("/register", logic.CreateUser)
	}

	user := r.Group("/")
	{
		user.GET("/getUser", logic.GetUser)
	}

	orders := r.Group("/")
	orders.Use(middleware.JWTAuth())
	{
		orders.GET("/getUserOrders", logic.GetUserOrders)
	}

	captcha := r.Group("/")
	{
		captcha.GET("/captcha", func(context *gin.Context) {
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

		captcha.POST("/captcha/verify", func(context *gin.Context) {
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
