package app

import (
	"zaiyun.app/app/model"
	"zaiyun.app/app/router"
)

func Start() {
	model.InitDB()
	defer model.CloseDB()
	router.InitRouter()
}
