package model

import "time"

func GetUserOrders(id int64) []Order {
	return nil
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
