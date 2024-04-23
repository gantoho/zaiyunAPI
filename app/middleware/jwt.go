package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"zaiyun.app/app/tools"
)

type MyClaims struct {
	UserID int64 `json:"user-id"`
	jwt.StandardClaims
}

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if gin.Mode() == gin.DebugMode {
			cookieToken, _ := context.Cookie("token")
			if cookieToken != "" {
				token = cookieToken
			}
		}
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
			context.JSON(http.StatusOK, tools.ECode{
				Code:    200,
				Message: "token 无效",
				Data:    nil,
			})
			context.Abort()
			return
		}
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
		return j.SigningKey, nil // 替换为实际密钥
	})

	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 如果 token 已过期，这里可以返回特定错误或处理逻辑
				return nil, errors.New("token 过期")
			}
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token 无效")
}
