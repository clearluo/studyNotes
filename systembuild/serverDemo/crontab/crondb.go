package crontab

import (
	"serverDemo/common/log"
	"time"
)

// 定时更新排行榜
func tickerUpdateDb() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("tickerUpdateDb panic:", err)
		}
	}()
	updateDb()
	dbTicker := time.NewTicker(time.Second * 60)

	for {
		select {
		case <-dbTicker.C:
			log.Info("定时更新静态数据缓存")
			updateDb()
		}
	}
}

func updateDb() {

}
