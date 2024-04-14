package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"zaiyun.app/app/models"
	"zaiyun.app/app/service"
	"zaiyun.app/app/tools"
)

type UUID struct {
	UUID string `json:"uuid"`
}

func (uuid UUID) String() string {
	return uuid.UUID
}

type MyClaims struct {
	Username string `json:"username"`
	UUID
	BufferTime int64 `json:"buffer-time"`
	jwt.StandardClaims
}

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		fmt.Printf("+++++++++++ %+v\n +++++++++++", token)

		if token == "" {
			context.JSON(http.StatusOK, tools.ECode{
				Code:    200,
				Message: "非法访问",
				Data:    nil,
			})
			context.Abort()
			return
		}

		if service.IsTokenExpired(token) {
			context.JSON(http.StatusOK, tools.ECode{
				Code:    200,
				Message: "Token已过期",
				Data:    nil,
			})
			context.Abort()
			return
		}

		j := NewJWT()

		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				context.JSON(http.StatusOK, tools.ECode{
					Code:    200,
					Message: "Token已过期",
					Data:    nil,
				})
				context.Abort()
				return
			}
			context.JSON(http.StatusOK, tools.ECode{
				Code:    200,
				Message: "Token无效",
				Data:    nil,
			})
			context.Abort()
			return
		}

		if err, _ = service.FindUserByUsername(claims.Username); err != nil {
			_ = service.JsonInBlacklist(models.JwtBlacklist{Jwt: token})
			context.JSON(http.StatusOK, tools.ECode{
				Code:    200,
				Message: "Token无效",
				Data:    nil,
			})
			context.Abort()
			return
		}

		if err, _ = service.FindUserByUuid(claims.UUID.String()); err != nil {
			_ = service.JsonInBlacklist(models.JwtBlacklist{Jwt: token})
			context.JSON(http.StatusOK, tools.ECode{
				Code:    200,
				Message: "Token无效",
				Data:    nil,
			})
			context.Abort()
			return
		}

		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + time.Now().Unix() + 60*60*2
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			context.Header("new-token", newToken)
			context.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
				err, RedisJwtToken := service.GetRedisJWT(newClaims.Username)
				if err != nil {
					//
				} else {
					_ = service.JsonInBlacklist(models.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = service.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		context.Set("claims", claims)
		context.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

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
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
