package retmsg

type msg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (m msg) Return() *message {
	return &message{
		Code: m.Code,
		Msg:  m.Msg,
	}
}

var (
	OK       = msg{200, "成功"}
	ERR_PARM = msg{401, "参数错误"}
	ERR_SYS  = msg{402, "系统错误"}
	ERR_DB   = msg{403, "数据库异常"}

	USER_LOGOUT     = msg{10001, "用户未登录"}
	USER_NO_EXIST   = msg{10002, "用户不存在"}
	USER_PASWD_FAIL = msg{10003, "用户密码不正确"}
	USER_ADD_FAIL   = msg{10004, "添加用户失败"}
	USER_ALREADY    = msg{10005, "用户已经存在"}

	NODE_ADD_ERR = msg{20001, "新增节点失败"}
)
