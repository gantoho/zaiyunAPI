package app

import (
	"zaiyun.app/app/models"
	"zaiyun.app/app/router"
	"zaiyun.app/app/schedule"
)

func Start() {
	models.InitDB()
	defer models.CloseDB()
	schedule.Start()
	router.InitRouter()
}
