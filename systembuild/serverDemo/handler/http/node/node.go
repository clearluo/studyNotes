package node

import (
	"net/http"
	"serverDemo/common/log"
	"serverDemo/common/retmsg"
	"serverDemo/db"
	"serverDemo/db/common"
	"serverDemo/db/node"

	"github.com/gin-gonic/gin"
)

func DoNode(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	action := c.Param("action")
	switch action {
	case "getNodeList":
		getNodeList(c)
	case "doNode":
		doNode(c)
	default:
		msg := retmsg.ERR_PARM.Return()
		log.Warn(msg.Msg)
		c.JSON(http.StatusOK, msg)
	}
}

func getNodeList(c *gin.Context) {
	msg := retmsg.OK.Return()
	defer func() {
		c.JSON(http.StatusOK, msg)
	}()
	rows, _ := node.SelectNodeList()
	msg.Data = rows
	return
}

func doNode(c *gin.Context) {
	msg := retmsg.OK.Return()
	defer func() {
		c.JSON(http.StatusOK, msg)
	}()
	req := &struct {
		Do          string `json:"do"` // add-增加；modify修改；del-删除
		Id          int    `json:"id"` // 修改和删除的时候必填；增加的时候不填
		NodeName    string `json:"nodeName"`
		HostIP      string `json:"hostIp"`
		Cluster     string `json:"cluster"`
		NodeType    string `json:"nodeType"`
		RpcUrl      string `json:"rpcUrl"`
		Frontend    int    `json:"frontend"`
		FrontendUrl string `json:"frontendUrl"`
		ConfigData  string `json:"configData"`
		GeoList     string `json:"geoList"`
	}{}
	c.BindJSON(req)
	switch req.Do {
	case "add":
		if len(req.NodeName) < 1 || len(req.HostIP) < 1 || len(req.Cluster) < 1 || len(req.NodeType) < 1 {
			msg = retmsg.ERR_PARM.Return()
			log.Warn(msg.Msg)
			return
		}
		if err := node.AddNode(req.NodeName, req.HostIP, req.Cluster, req.NodeType, req.RpcUrl, req.Frontend,
			req.FrontendUrl, req.ConfigData, req.GeoList); err != nil {
			log.Warn(err)
			msg = retmsg.NODE_ADD_ERR.Return()
			log.Warn(msg.Msg)
			return
		}
	case "modify":
		if len(req.NodeName) < 1 || len(req.HostIP) < 1 || len(req.Cluster) < 1 || len(req.NodeType) < 1 {
			msg = retmsg.ERR_PARM.Return()
			log.Warn(msg.Msg)
			return
		}
		updateMap := map[string]interface{}{
			"id":     req.Id,
			"name":   req.NodeName,
			"hostIp": req.HostIP,
		}
		if err := common.UpdateTableById(db.GetAppDb(), &node.TNodeInfo{}, updateMap); err != nil {
			log.Warn(err)
		}

	case "del":
		if req.Id < 1 {
			msg = retmsg.ERR_PARM.Return()
			log.Warn(msg.Msg)
			return
		}
		if err := node.DelNode(req.Id); err != nil {
			log.Warn(err)
		}
	default:
		msg = retmsg.ERR_PARM.Return()
		log.Warn(msg.Msg)
		return
	}
	msg.Data, _ = node.SelectNodeList()
	return
}
