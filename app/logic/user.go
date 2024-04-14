package logic

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
	"time"
	"zaiyun.app/app/middleware"
	"zaiyun.app/app/models"
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
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "server: login error"})
		return
	}
	ret, err := models.PostLogin(user.Username)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "login error"})
		return
	}
	if ret.Password != encrypt(user.Password) {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "密码错误"})
		return
	}
	context.SetCookie("id", strconv.FormatInt(ret.ID, 10), 86400, "/", "localhost", false, false)
	newJwt := middleware.NewJWT()
	token, _ := newJwt.CreateToken(middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    user.Username,
			NotBefore: time.Now().Unix() - 60,
		},
	})
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": gin.H{"id": ret.ID, "username": ret.Username, "token": token, "avatar": ret.Avatar, "motto": ret.Motto, "created_time": ret.CreatedTime.Format("2006-01-02 15:04:05"), "updated_time": ret.UpdatedTime.Format("2006-01-02 15:04:05")}, "message": "login success"})
}

type CUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

func CreateUser(context *gin.Context) {
	var user CUser
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "server: create user error"})
		return
	}

	//fmt.Printf("%+v\n", user)
	//
	//encrypt(user.Password)
	//return

	if user.Username == "" || user.Password == "" || user.Password2 == "" {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "必填项不能为空"})
		return
	}

	if user.Password != user.Password2 {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "两次密码不同"})
		return
	}

	//usernameLen := len(user.Username)
	//passwordLen := len(user.Password)
	//if usernameLen < 3 || usernameLen > 18 || passwordLen < 3 || passwordLen > 18 {
	//	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "用户名长度为3-18"})
	//	return
	//}

	regexUser := regexp.MustCompile("^[a-zA-Z0-9_-]{3,18}$")
	if !regexUser.MatchString(user.Username) {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "用户名只能包含字母、数字、下划线、横线, 长度大于3小于18"})
		return
	}

	passwordRegex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[-_]).{6,20}$`
	regexPass, _ := regexp2.Compile(passwordRegex, 0)
	m, err := regexPass.FindStringMatch(user.Password)
	fmt.Print(m, err)
	if err != nil {
		return
	}
	if m == nil {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "密码必须由包含大写字母，小写字母，数字，下划线_，短横线-，且长度大于6小于30"})
		return
	}

	var _getUser models.User
	_getUser, _ = models.GetUser(user.Username)
	if _getUser.ID > 0 {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "用户已存在"})
		return
	}

	newUser := models.User{
		Username:    user.Username,
		Password:    encrypt(user.Password),
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	err = models.CreateUser(&newUser)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "message": "server: create user error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": gin.H{"id": newUser.ID, "username": newUser.Username, "avatar": newUser.Avatar, "motto": newUser.Motto, "created_time": newUser.CreatedTime.Format("2006-01-02 15:04:05"), "updated_time": newUser.UpdatedTime.Format("2006-01-02 15:04:05")}, "message": "注册成功"})
}

func encrypt(pwd string) string {
	nawPwd := pwd + "zaiyunapp123"
	hash := md5.New()
	hash.Write([]byte(nawPwd))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
