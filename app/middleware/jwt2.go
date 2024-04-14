package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func JWTAuth2() {
	// 签发
	mySignedString := []byte("AllYourBase")

	// 第二种写法 MapClaims map
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "gantoho",
		"exp":      time.Now().Unix() + 60*60*2,
		"iss":      "gantoho",
		"nbf":      time.Now().Unix() - 60,
	})

	ss, err := token.SignedString(mySignedString)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		panic(err)
	}
	fmt.Printf("token: %+v\n", token)
	fmt.Printf("ss: %+v\n", ss)

	// 解密
	token, err = jwt.ParseWithClaims(ss, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySignedString, nil
	})
	if err != nil {
		fmt.Printf("err: %s\n", err)
		panic(err)
	}
	fmt.Println(token)
	fmt.Println(token.Claims.(*jwt.MapClaims))
}
