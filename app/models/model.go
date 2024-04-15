package models

import "time"

// Goods undefined
type Goods struct {
	ID          int64     `json:"id" gorm:"id"`
	Title       string    `json:"title" gorm:"title"`
	Price       float64   `json:"price" gorm:"price"`
	Cover       int64     `json:"cover" gorm:"cover"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*Goods) TableName() string {
	return "goods"
}

// Images undefined
type Images struct {
	ID          int64     `json:"id" gorm:"id"`
	URL         string    `json:"url" gorm:"url"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*Images) TableName() string {
	return "images"
}

// SwiperList undefined
type SwiperList struct {
	ID          int64     `json:"id" gorm:"id"`
	GoodsId     int64     `json:"goods_id" gorm:"goods_id"`
	ImagesId    int64     `json:"images_id" gorm:"images_id"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*SwiperList) TableName() string {
	return "swiper_list"
}

// IndexSwiper undefined
type IndexSwiper struct {
	ID          int64     `json:"id" gorm:"id"`
	URL         string    `json:"url" gorm:"url"`
	To          string    `json:"to" gorm:"to"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*IndexSwiper) TableName() string {
	return "index_swiper"
}

// User undefined
type User struct {
	ID          int64     `json:"id" gorm:"id"`
	Username    string    `json:"username" gorm:"username"`
	Password    string    `json:"-" gorm:"password"`
	Avatar      string    `json:"avatar" gorm:"avatar"`
	Motto       string    `json:"motto" gorm:"motto"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*User) TableName() string {
	return "user"
}

// Order undefined
type Order struct {
	ID          int64     `json:"id" gorm:"id"`
	UserId      int64     `json:"user_id" gorm:"user_id"`
	GoodsId     int64     `json:"goods_id" gorm:"goods_id"`
	Status      string    `json:"status" gorm:"status"` // 0 表示下单未付款\r\n1 表示订单已失效\r\n2 表示订单已付款\r\n3 表示订单已完成\r\n4 表示订单存在异常
	Address     string    `json:"address" gorm:"address"`
	Expire      int64     `json:"expire" gorm:"expire"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*Order) TableName() string {
	return "order"
}
