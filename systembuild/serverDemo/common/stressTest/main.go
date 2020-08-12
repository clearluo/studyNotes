package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const (
	goN        int = 10000
	timeSecond int = 300
)

type RetData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func httpReq() bool {
	client := &http.Client{}
	data := make(map[string]interface{})
	data["name"] = "zhaofan"
	data["pwd"] = "23"
	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "http://192.168.9.204/serverDemo/test/abc", bytes.NewReader(bytesData))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	body, _ := ioutil.ReadAll(resp.Body)
	ret := RetData{}
	if err := json.Unmarshal(body, &ret); err != nil {
		fmt.Println(err, ": ", string(body))
		return false
	}
	if ret.Code == 200 {
		return true
	}
	return false
}

func main() {
	start := time.Now()
	allFail := make([]int, goN)
	allSucc := make([]int, goN)
	w := sync.WaitGroup{}
	for i := 0; i < goN; i++ {
		go func(index int) {
			w.Add(1)
			defer w.Done()
			for j := 0; j < timeSecond; j++ {
				if httpReq() {
					allSucc[index]++
				} else {
					allFail[index]++

				}
				time.Sleep(time.Second)
			}
		}(i)
	}
	w.Wait()
	failN, succN := 0, 0
	for _, v := range allFail {
		failN += v
	}
	for _, v := range allSucc {
		succN += v
	}
	sumReq := failN + succN
	rate := float64(succN) / float64(sumReq)
	use := time.Since(start)
	fmt.Printf("%d并发请求持续%ds,总共测试:%d 失败:%d 成功率:%.2f%% 用时:%v 每次请求耗时:%.2fms\n",
		goN, timeSecond, sumReq, failN, rate*100, use, float64(use.Milliseconds())/float64(sumReq))
}
