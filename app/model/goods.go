package model

import "fmt"

type GoodsLinkImages struct {
	Goods
	Images
}

// 第一个参数是页码pageCode，第二个参数是每页数量pageSize
func GetGoods(args ...int) []map[string]interface{} {
	if len(args) == 0 {
		args = append(args, 1)
		args = append(args, 10)
	}
	if len(args) == 1 {
		args = append(args, 10)
	}
	var goods []map[string]interface{}
	err := Conn.Table("goods").
		Offset((args[0]-1)*args[1]).
		Limit(args[1]).
		Model(&GoodsLinkImages{}).
		Select("goods.*", "images.url").
		Joins("left join images on goods.cover = images.id").Find(&goods).Error
	if err != nil {
		panic(err)
	}
	return goods
}

type GoodsLinkSwiperList struct {
	Goods
	Swipers []string `json:"swipers" gorm:"swipers"`
}

func GetGoodsDetail(id int64) GoodsLinkSwiperList {
	var goods Goods
	err := Conn.Table("goods").
		Where("goods.id = ?", id).First(&goods).Error
	if err != nil {
		fmt.Printf("err: %v", err)
		panic(err)
	}
	fmt.Printf("goods: %+v", goods)
	var swipers []string
	err = Conn.Table("swiper_list").
		Select("images.url").
		Where("swiper_list.goods_id = ?", id).
		Joins("left join images on images.id = swiper_list.images_id").
		Find(&swipers).Error
	if err != nil {
		fmt.Printf("err: %v", err)
		panic(err)
	}
	var goodsLinkSwiperList GoodsLinkSwiperList
	goodsLinkSwiperList.Goods = goods
	goodsLinkSwiperList.Swipers = swipers
	return goodsLinkSwiperList
}
