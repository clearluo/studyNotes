package runbefore

import (
	"serverDemo/common/basic"
	"serverDemo/common/myredis"
	"serverDemo/crontab"
	"serverDemo/db"
)

// InitRun init
func InitRun() {
	// 初始化Msyql
	db.InitDb()
	// 加载模板表到map
	//tpl.InitTplData()
	// 数据库相关数据加载入Map
	//db.InitMap()
	// 初始化Redis
	myredis.InitRedis()
	// 初始化敏感词
	// trie.InitTrie()
	// 初始化GeoIp2
	//geoip.InitGeoIp()
	// 初始化定时任务
	if basic.App.RunCron {
		go func() {
			crontab.CronMain()
		}()
	}
}
