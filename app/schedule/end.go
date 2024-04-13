package schedule

import (
	"time"
	"zaiyun.app/app/models"
)

func Start() {
	go OrderEnd()
	return
}

func OrderEnd() {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			//fmt.Printf("定时器 voteEnd 启动")
			_ = models.OrderEnd()
		}
	}
}
