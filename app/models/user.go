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

func GetUserByUserName(username string) (User, error) {
	var user User
	err := Conn.Table("user").Where("username = ?", username).First(&user).Error
	return user, err
}

func GetUserByUserID(id int64) (User, error) {
	var user User
	err := Conn.Table("user").Where("id = ?", id).First(&user).Error
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
