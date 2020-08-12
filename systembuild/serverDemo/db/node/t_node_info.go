package node

import (
	"serverDemo/db"
	"time"
)

type TNodeInfo struct {
	Id          int       `xorm:"name id not null pk autoincr INT(11)" json:"id"`
	Name        string    `xorm:"name name not null default '' comment('节点node名称') index VARCHAR(50)" json:"name"`
	HostId      string    `xorm:"name hostId default '' comment('在哪台机器上运行') VARCHAR(255)" json:"hostId"`
	Cluster     string    `xorm:"name cluster default '' comment('所在的集群') VARCHAR(255)" json:"cluster"`
	NodeType    string    `xorm:"name nodeType default '' comment('服务的类型') VARCHAR(255)" json:"nodeType"`
	IsOnline    int       `xorm:"name isOnline default 0 TINYINT(1)" json:"isOnline"`
	RpcUrl      string    `xorm:"name rpcUrl default '' comment('对集群开放的ip端口') VARCHAR(255)" json:"rpcUrl"`
	Frontend    int       `xorm:"name frontend default 0 comment('是否是网关的节点') TINYINT(1)" json:"frontend"`
	FrontendUrl string    `xorm:"name frontendUrl not null default '' comment('网关链接的地址wss://h5tomb.99.com:3011') VARCHAR(64)" json:"frontendUrl"`
	LastTime    time.Time `xorm:"name lastTime not null default CURRENT_TIMESTAMP comment('节点心跳时间') TIMESTAMP" json:"lastTime"`
	ConfigData  string    `xorm:"name configData comment('启动的配置文件') TEXT" json:"configData"`
	GeoList     string    `xorm:"name geoList default '' comment('ip白名单，非法ip不允许获取配置') VARCHAR(255)" json:"geoList"`
	CreateAt    time.Time `xorm:"name createAt not null default CURRENT_TIMESTAMP TIMESTAMP" json:"createAt"`
}

func (t *TNodeInfo) TableName() string {
	return "t_node_info"
}

func AddNode(nodeName string, hostIp string, cluster string, nodeType string, rpcUrl string, frontend int,
	frontendUrl string, configData string, geoList string) (err error) {
	sql := `INSERT INTO t_node_info(name,hostId,cluster,nodeType,rpcUrl,frontend,frontendUrl,configData,geoList)VALUES
		(?,?,?,?,?,?,?,?,?)`
	_, err = db.GetAppDb().Exec(sql, nodeName, hostIp, cluster, nodeType, rpcUrl, frontend, frontendUrl, configData, geoList)
	return
}

func DelNode(id int) (err error) {
	sql := `DELETE FROM t_node_info WHERE id = ?`
	_, err = db.GetAppDb().Exec(sql, id)
	return
}

func SelectNodeList() ([]*TNodeInfo, error) {
	rows := []*TNodeInfo{}
	sql := `SELECT * FROM t_node_info`
	db.GetAppDb().SQL(sql).Find(&rows)
	return rows, nil
}
