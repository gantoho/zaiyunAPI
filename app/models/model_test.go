package models

import (
	"fmt"
	"testing"
)

func TestGetGoods(t *testing.T) {
	InitDB()
	defer CloseDB()

	goods := GetGoods()
	fmt.Printf("%+v\n", goods)
}

func TestGetGoodsDetail(t *testing.T) {
	InitDB()
	defer CloseDB()

	goods := GetGoodsDetail(1)
	fmt.Printf("%+v\n", goods)
}

func TestGetIndexSwipers(t *testing.T) {
	InitDB()
	defer CloseDB()

	swipers := GetIndexSwipers()
	fmt.Printf("%+v\n", swipers)
}
