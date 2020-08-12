package myredis

import (
	"fmt"
	"time"
)

// 获取实时上报信息的key
func GetMemKey(action string, node string) string {
	return fmt.Sprintf("mem:%v:%v", action, node)
}

// 获取节点信息的key
func GetNodeKey(node string) string {
	return fmt.Sprintf("node:%v", node)
}

// 获取monitor qps的key
func GetMonitorQpsKey(minute int) string {
	return fmt.Sprintf("qps:monitor:%v", minute)
}

type NodeData struct {
	Node     string    `json:"node"`
	LastTime time.Time `json:"lastTime"`
	State    int       `json:"state"` // 0-正常，1-异常，2-离线
}
