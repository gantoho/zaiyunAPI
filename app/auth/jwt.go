package auth

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID   int64
	Username string
	jwt.RegisteredClaims
}
