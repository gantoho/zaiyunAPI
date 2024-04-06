package model

import "fmt"

func PostLogin(username, password string) (User, error) {
	var user User
	err := Conn.Table("user").Where("username = ? and password = ?", username, password).First(&user).Error
	if err != nil {
		fmt.Printf("login error: %v", err)
	}
	return user, err
}
