package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewJwt() {
	// 签发
	var mc = MyClaims{
		Username: "gantoho",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "gantoho",
			NotBefore: time.Now().Unix() - 60,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	mySignedString := []byte("AllYourBase")
	ss, err := token.SignedString(mySignedString)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		panic(err)
	}
	fmt.Printf("token: %+v\n", token)
	fmt.Printf("ss: %+v\n", ss)

	// 解密
	token, err = jwt.ParseWithClaims(ss, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySignedString, nil
	})
	if err != nil {
		fmt.Printf("err: %s\n", err)
		panic(err)
	}
	fmt.Println(token)
	fmt.Println(token.Claims.(*MyClaims).Username)
}
