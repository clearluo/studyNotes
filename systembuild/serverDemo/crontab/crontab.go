package crontab

import (
	"serverDemo/common/log"
	"time"
)

var funcMap map[string]func(int) error

func init() {
	funcMap = make(map[string]func(int) error)
}

func CronMain() {
	log.Info("CronMain")
	defer func() {
		if err := recover(); err != nil {
			log.Error("CronMain recover:", err)
		}
	}()
	//go tickerUpdateDb()
	go tickerUpdateDbFromRedis()
	ticker := time.NewTicker(3600 * time.Second)
	for {
		select {
		case <-ticker.C:
			execCronTask()
		}
	}
}

func execCronTask() {
	//util.GetFilterList()
}
