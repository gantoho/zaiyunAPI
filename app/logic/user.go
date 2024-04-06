package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zaiyun.app/app/model"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func PostLogin(context *gin.Context) {
	var user User
	err := context.ShouldBind(&user)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	ret, err := model.PostLogin(user.Username, user.Password)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "login error"})
		return
	}
	context.SetCookie("id", strconv.FormatInt(ret.ID, 10), 86400, "/", "localhost", false, false)
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": ret, "message": "login success"})
}
