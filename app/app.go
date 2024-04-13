package app

import (
	"zaiyun.app/app/model"
	"zaiyun.app/app/router"
	"zaiyun.app/app/schedule"
)

func Start() {
	model.InitDB()
	defer model.CloseDB()
	schedule.Start()
	router.InitRouter()
}
