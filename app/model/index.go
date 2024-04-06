package model

import "fmt"

func GetIndexSwipers() []IndexSwiper {
	var indexSwipers []IndexSwiper
	err := Conn.Table("index_swiper").Find(&indexSwipers).Error
	if err != nil {
		fmt.Printf("GetIndexSwipers err: %+v", err)
	}
	return indexSwipers
}
