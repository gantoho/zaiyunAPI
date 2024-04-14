package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"zaiyun.app/app/tools"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")

		if token == "" {
			context.JSON(http.StatusOK, tools.ECode{
				Code:    200,
				Message: "非法访问",
				Data:    nil,
			})
			context.Abort()
		}

		return

		//// 签发
		//mySignedString := []byte("AllYourBase")
		//
		//// 第一种写法 StandardClaims struct
		//var mc = MyClaims{
		//	Username: "gantoho",
		//	StandardClaims: jwt.StandardClaims{
		//		ExpiresAt: time.Now().Unix() + 60*60*2,
		//		Issuer:    "gantoho",
		//		NotBefore: time.Now().Unix() - 60,
		//	},
		//}
		//token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
		//
		//ss, err := token.SignedString(mySignedString)
		//if err != nil {
		//	fmt.Printf("err: %s\n", err)
		//	panic(err)
		//}
		//fmt.Printf("token: %+v\n", token)
		//fmt.Printf("ss: %+v\n", ss)
		//
		//// 解密
		//token, err = jwt.ParseWithClaims(ss, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		//	return mySignedString, nil
		//})
		//if err != nil {
		//	fmt.Printf("err: %s\n", err)
		//	panic(err)
		//}
		//fmt.Println(token)
		//fmt.Println(token.Claims.(*MyClaims).Username)
	}
}
