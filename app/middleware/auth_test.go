package middleware

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestCreate(t *testing.T) {
	newJwt := NewJWT()
	mc := MyClaims{
		Username: "",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "gantoho",
			NotBefore: time.Now().Unix() - 60,
		},
	}
	token, err := newJwt.CreateToken(mc)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Printf("%+v\n", token)
}
