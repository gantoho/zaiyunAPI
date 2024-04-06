package model

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
	Password    string    `json:"-" gorm:"-"`
	Avatar      string    `json:"avatar" gorm:"avatar"`
	Motto       string    `json:"motto" gorm:"motto"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*User) TableName() string {
	return "user"
}
