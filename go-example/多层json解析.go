package main

import (
	"encoding/json"
	"fmt"
)

type Item1 struct {
	Errno  int     `json:"errno"`
	Errmsg string  `json:"errmsg"`
	Data   []Item2 `json:"data"`
}
type Item2 struct {
	Price         string `json:"price"`
	ProcessId     string `json:"process_id"`
	ProcessName   string `json:"process_name"`
	ProcessDefine Item3  `json:"process_define"`
}
type Item3 struct {
	TimeLimit int     `json:"time_limit"`
	Steps     []Item4 `json:"steps"`
}
type Item4 struct {
	StepName      string `json:"step_name"`
	ProcessStepId int    `json:"process_step_id"`
	IsNeedReview  int    `json:"is_need_review"`
	TimeLimit     int    `json:"time_limit"`
}

func main() {
	str := `{"errno":0,"errmsg":"成功","data":[{"price":"40000","process_define":{"steps":[{"is_need_review":1,"process_step_id":162,"step_name":"工商注册","step_order_seq":1,"time_limit":3},{"is_need_review":1,"process_step_id":163,"step_name":"基本开户","step_order_seq":2,"time_limit":5}],"time_limit":8},"process_id":"16","process_name":"公司注册"},{"price":"36000","process_define":{"time_limit":30,"steps":[{"step_name":"接票","process_step_id":164,"is_need_review":1,"step_order_seq":1,"time_limit":3},{"step_name":"做帐","process_step_id":165,"is_need_review":1,"step_order_seq":2,"time_limit":15},{"step_name":"报税","process_step_id":166,"is_need_review":1,"step_order_seq":3,"time_limit":12}]},"process_id":"17","process_name":"代理记账"}]}`
	var o Item1
	if err := json.Unmarshal([]byte(str), &o); err != nil {
		fmt.Println(err)
	}
	fmt.Println(o)

	var m interface{}
	if err := json.Unmarshal([]byte(str), &m); err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)

}
