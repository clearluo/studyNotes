package test

import (
	"net/http"
	"serverDemo/common/log"
	"serverDemo/common/retmsg"

	"github.com/gin-gonic/gin"
)

func DoTest(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	action := c.Param("action")
	switch action {
	case "abc":
		abc(c)
	default:
		msg := retmsg.ERR_PARM.Return()
		log.Warn(msg.Msg)
		c.JSON(http.StatusOK, msg)
	}
}

func abc(c *gin.Context) {
	msg := retmsg.OK.Return()
	defer func() {
		c.JSON(http.StatusOK, msg)
	}()
	req := &struct {
		Name string `json:"name"`
		Pwd  string `json:"pwd"`
	}{}
	c.BindJSON(req)
	//log.Infof("[req]%v:%v", util.RunFuncName(), util.AssertMarshal(req))
	return
}
