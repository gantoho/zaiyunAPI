package models

import "fmt"

func PostLogin(username string) (User, error) {
	var user User
	err := Conn.Table("user").Where("username = ?", username).First(&user).Error
	if err != nil {
		fmt.Printf("login error: %v", err)
	}
	return user, err
}

func GetUser(username string) (User, error) {
	var user User
	err := Conn.Table("user").Where("username = ?", username).First(&user).Error
	return user, err
}

func CreateUser(user *User) error {
	fmt.Printf("user: %+v ****************", user)
	err := Conn.Table("user").Create(&user).Error
	if err != nil {
		fmt.Printf("create user error: %v", err)
		return err
	}
	return nil
}
