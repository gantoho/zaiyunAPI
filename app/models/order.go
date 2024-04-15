package models

import (
	"fmt"
	"time"
)

func GetUserOrders(id int64, args ...int) ([]Order, error) {
	if len(args) == 0 {
		args = append(args, 1)
		args = append(args, 10)
	}
	if len(args) == 1 {
		args = append(args, 10)
	}
	orders := make([]Order, 0)
	err := Conn.Table("order").
		Offset((args[0]-1)*args[1]).
		Limit(args[1]).
		Where("user_id = ?", id).Find(&orders).Error
	if err != nil {
		fmt.Printf("GetUserOrders Error: %+v", err)
	}
	return orders, err
}

func OrderEnd() error {
	orders := make([]Order, 0)
	err := Conn.Table("order").Where("status = ?", 0).Find(&orders).Error
	if err != nil {
		return err
	}
	for _, order := range orders {
		if order.Expire+order.CreatedTime.Unix() <= time.Now().Unix() {
			err := Conn.Table("order").Where("id = ?", order.ID).Update("status", 1).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
