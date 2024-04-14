package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"zaiyun.app/app/tools"
)

type MyClaims struct {
	Username   string `json:"username"`
	BufferTime int64  `json:"buffer-time"`
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
			return
		}

		j := NewJWT()

		claims, err := j.ParseToken(token)

		if err != nil {
			fmt.Printf("JWTAuth Error: %+v\n", err)
			context.JSON(http.StatusOK, tools.ECode{
				Code:    200,
				Message: "token失效",
				Data:    nil,
			})
			context.Abort()
			return
		}
		fmt.Printf("JWTAuth claims: %+v\n", claims)
		context.Set("claims", claims)
		context.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("jwt token key"),
	}
}

// 创建一个token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析token
func (j *JWT) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	fmt.Printf("ParseToken: %+v +++++++++++ %s\n", token, err)

	if err != nil {
		fmt.Printf("ParseToken Error: %+v +++++++++++ %s\n", token, err)
		return nil, err
	}
	claims, _ := token.Claims.(*MyClaims)

	return claims, err
}
