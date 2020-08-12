package crontab

import (
	"serverDemo/common/log"
	"time"
)

// 定时降redis中的数据落地到db
func tickerUpdateDbFromRedis() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("tickerUpdateDbFromRedis panic:", err)
		}
	}()
	updateDbFromRedis()
	dbTicker := time.NewTicker(time.Second * 60)

	for {
		select {
		case <-dbTicker.C:
			log.Info("定时更新redis数据")
			updateDbFromRedis()
		}
	}
}

func updateDbFromRedis() {
	memHandler()
	qpsHandler()
}
func memHandler() {
	log.Info("memHandler")

}

func qpsHandler() {
	log.Info("qpsHandler")
}
