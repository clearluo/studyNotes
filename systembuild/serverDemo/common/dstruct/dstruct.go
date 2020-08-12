package dstruct

import "time"

type ReqHead struct {
	//AppID     string `json:"app_id"`
	//Secret    string `json:"secret"`
	Version   string `form:"version"`
	Uid       int    `form:"uid"`
	Timestamp int64  `form:"timestamp"`
	Method    string `form:"method"`
	Sign      string `form:"sign"`
}
type Author struct {
	Username string `json:"name"`
	Token    string `json:"token"`
	Role     string `json:"role"`
	UserId   int    `json:"userId"`
	Expired  int64  `json:"expired"`
}

type TokenDsta struct {
	UserId      int   `json:"userId"`
	MilliSecond int64 `json:"milliSecond"`
	AreaId      int   `json:"areaId"`
	AuthFlg     int   `json:"authFlg"`
}

type AdminList struct {
	Id       int       `orm:"column(id)" json:"id"`
	Username string    `orm:"column(username)" json:"username"`
	Role     string    `orm:"column(role)" json:"role"`
	OnUse    int8      `orm:"column(onUse)" json:"onUse"`
	UpdateAt time.Time `orm:"column(updateAt);type(timestamp) json:"updateAt"`
}

type Server struct {
	Server          interface{} `json:"server"`
	Net             ServerNet   `json:"net"`
	GamePlatform    int         `json:"gamePlatform"`
	Path            ServerPath  `json:"path"`
	IsDebugVersions bool        `json:"isDebugVersions"`
	Tpl             ServerTpl   `json:"tpl"`
	Opendns         []string    `json:"opendns"`
	Token           string      `json:"token"`
	RealAuth        int         `json:"realAuth"`
}
type ServerNet struct {
	BaseUrl     string `json:"baseUrl"`
	LoginUrl    string `json:"loginUrl"`
	RegUrl      string `json:"regUrl"`
	GateHost    string `json:"gateHost"`
	GatePort    string `json:"gatePort"`
	HttpRequest string `json:"httpRequest"`
}
type ServerPath struct {
	Res  string `json:"res"`
	Anim string `json:"anim"`
	Data string `json:"data"`
}
type ServerTpl struct {
	Path string `json:"path"`
}

type ServerStatus struct {
	ServerName string                 `json:"serverName"` // 服务器名字
	MsgType    string                 `json:"msgType"`    // 消息类型：心跳、资源状态、当前压力等
	Data       map[string]interface{} `json:"data"`       // 根据不同消息类型，对应不同结构
}

type AddUserNum struct {
	Date     string `xorm:"date" json:"date"`
	UserType string `xorm:"userType" json:"userType"`
	Num      int    `xorm:"num" json:"num"`
}
type DayAddUserNum struct {
	Date string `json:"date"`
	All  int    `json:"all"`
	U99  int    `json:"u99"`
	Wx   int    `json:"wx"`
}

type TotalCount struct {
	UserType string `xorm:"userType" json:"userType"`
	Num      int    `xorm:"num" json:"num"`
}

type ServerDetail struct {
	Cpu []CpuDetail `json:"cpu"`
	Mem MemDetail   `json:"mem"`
}
type CpuDetail struct {
	Cpu       string  `json:"cpu"`
	User      float64 `json:"user"`
	System    float64 `json:"system"`
	Idle      float64 `json:"idle"`
	Nice      float64 `json:"nice"`
	Iowait    float64 `json:"iowait"`
	Irq       float64 `json:"irq"`
	Softirq   float64 `json:"softirq"`
	Steal     float64 `json:"steal"`
	Guest     float64 `json:"guest"`
	GuestNice float64 `json:"guestNice"`
}
type MemDetail struct {
	Total          int `json:"total"`
	Available      int `json:"available"`
	Used           int `json:"used"`
	UsedPercent    int `json:"usedPercent"`
	Free           int `json:"free"`
	Active         int `json:"active"`
	inactive       int `json:"inactive"`
	Wired          int `json:"wired"`
	Laundry        int `json:"laundry"`
	Buffers        int `json:"buffers"`
	Cached         int `json:"cached"`
	Writeback      int `json:"writeback"`
	Dirty          int `json:"dirty"`
	Writebacktmp   int `json:"writebacktmp"`
	Shared         int `json:"shared"`
	Slab           int `json:"slab"`
	Sreclaimable   int `json:"sreclaimable"`
	Sunreclaim     int `json:"sunreclaim"`
	Pagetables     int `json:"pagetables"`
	Swapcached     int `json:"swapcached"`
	Commitlimit    int `json:"commitlimit"`
	Committedas    int `json:"committedas"`
	Hightotal      int `json:"hightotal"`
	Highfree       int `json:"highfree"`
	Lowtotal       int `json:"lowtotal"`
	Lowfree        int `json:"lowfree"`
	Swaptotal      int `json:"swaptotal"`
	Swapfree       int `json:"swapfree"`
	Mapped         int `json:"mapped"`
	Vmalloctotal   int `json:"vmalloctotal"`
	Vmallocused    int `json:"vmallocused"`
	Vmallocchunk   int `json:"vmallocchunk"`
	Hugepagestotal int `json:"hugepagestotal"`
	Hugepagesfree  int `json:"hugepagesfree"`
	Hugepagesize   int `json:"hugepagesize"`
}
