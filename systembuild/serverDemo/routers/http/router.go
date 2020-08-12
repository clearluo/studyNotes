package http

import (
	"fmt"
	"net/http"
	"serverDemo/common/basic"
	"serverDemo/common/log"
	"serverDemo/common/myredis"
	"serverDemo/common/util"
	"serverDemo/handler/http/admin"
	"serverDemo/handler/http/node"
	"serverDemo/handler/http/test"
	"strings"

	"github.com/gin-gonic/gin"
	sessions "github.com/tommy351/gin-sessions"
)

// HistoryLog 修改历史log
var HistoryLog = []string{
	"serverDemo-backend_V1.0.0_2020-07-06_13:06_初始化",
}

// Version 当前版本号
var Version = "serverDemo-backend_V1.0.0_2020-07-06_13:06_初始化"

// Cors 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			//c.Header("Access-Control-Allow-Origin", "*")
			//c.Header("Access-Control-Allow-Methods", "*")
			//c.Header("Access-Control-Allow-Headers", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, token, userId")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
func BeforeFun() gin.HandlerFunc {
	return func(c *gin.Context) {
		myredis.GetDb().Incr(myredis.GetMonitorQpsKey(util.GetTimeByYyyymmddhhmm()))
	}
}

// InitRoute 初始化路由
func InitRoute() {
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("__session"))
	r.Use(sessions.Middleware("do_sesion", store))
	r.Use(Cors())
	r.Use(BeforeFun())
	r.POST("/serverDemo/admin/:action", admin.DoAdmin) //
	r.POST("/serverDemo/node/:action", node.DoNode)    //

	// Test
	r.POST("/serverDemo/test/:action", test.DoTest)

	log.Info("strat http server @port=", basic.App.Port)
	fmt.Println("strat http server @port=", basic.App.Port)
	log.Infof("History Log\n%v", strings.Join(HistoryLog, "\n"))
	log.Info("Version:", Version)
	if err := r.Run(":" + basic.App.Port); err != nil {
		panic("r.Run err: " + err.Error())
	}
}
