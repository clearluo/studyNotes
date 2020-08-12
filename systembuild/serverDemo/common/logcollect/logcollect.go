package main

import (
	"encoding/json"
	"fmt"
	"net"
	"serverDemo/common/log"
	"sync"
	"time"
)

var logBuf chan string
var wg sync.WaitGroup

//var conn net.Conn
var server = "127.0.0.1:1024"

type Event struct {
	TrackingId     string `json:"trackingId"` // 项目id，可以使用立项id "GMS0015"
	EventName      string `json:"eventName"`  // 时间名称，如 login
	EventPayload   User   `json:"eventPayload"`
	EventCreatedAt int64  `json:"eventCreatedAt"` // 毫秒时间戳
}
type User struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
}

func init() {
	logBuf = make(chan string, 1024)
}
func main() {
	go link()
	//wg.Add(100)
	for i := 0; i < 100; i++ {
		data := Event{
			TrackingId: "GMS0015",
			EventName:  "login",
			EventPayload: User{
				Id:   i,
				Data: "clearluo",
			},
			EventCreatedAt: time.Now().UnixNano() / 1000000,
		}
		b, _ := json.Marshal(data)
		Send(string(b) + "\n")
		time.Sleep(time.Second)
	}
	//wg.Wait()
	select {}
}

//负责连接以及连接恢复
func link() {
	for {
		conn, err := net.Dial("tcp", server)
		if err != nil {
			fmt.Print("connect fail")
		} else {
			fmt.Println("connect ok")
			doSend(conn)
		}
		time.Sleep(3 * time.Second)
	}
}

// 发送
func doSend(conn net.Conn) {
	defer conn.Close()
	//ticker := time.NewTicker(time.Second)
	for {
		select {
		case content := <-logBuf:
			if _, err := conn.Write([]byte(content)); err != nil {
				logBuf <- content
				log.Warn("send log to logcollect err:", err)
				return
			}
		}
	}
}

func Send(str string) bool {
	if len(str) < 1 {
		log.Warn("str is too short")
		return false
	}
	select {
	case logBuf <- str:
		return true
	default:
		// 写入文件
		log.Info("logcollect ", str)
	}
	return true
}
