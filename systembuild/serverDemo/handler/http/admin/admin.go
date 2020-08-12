package admin

import (
	"fmt"
	"net/http"
	"serverDemo/common/auth"
	"serverDemo/common/consts"
	"serverDemo/common/dstruct"
	"serverDemo/common/log"
	"serverDemo/common/retmsg"
	"serverDemo/common/util"
	"serverDemo/common/util/session"
	"serverDemo/db"
	"serverDemo/db/admin"
	"serverDemo/db/common"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// DoLogin 登录
func DoAdmin(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	action := c.Param("action")
	switch action {
	case "login":
		login(c)
	case "logout":
		logout(c)
	case "changePwd":
		changePwd(c)
	case "addUser":
		addUser(c)
	case "modifyUser":
		modifyUser(c)
	case "listUser":
		listUser(c)
	default:
		msg := retmsg.ERR_PARM.Return()
		log.Warn(msg.Msg)
		c.JSON(http.StatusOK, msg)
	}
}

func login(c *gin.Context) {
	msg := retmsg.OK.Return()
	defer func() {
		c.JSON(http.StatusOK, msg)
	}()
	req := &struct {
		Name string `json:"name"`
		Pwd  string `json:"pwd"`
	}{}
	c.BindJSON(req)
	log.Info("req:", util.AssertMarshal(req))

	if len(req.Name) < 2 || len(req.Pwd) != 32 {
		msg = retmsg.ERR_PARM.Return()
		log.Warn(msg.Msg)
		return
	}
	adminRow, err := admin.SelectAdminByName(req.Name)
	if err != nil {
		msg = retmsg.USER_NO_EXIST.Return()
		log.Warn(msg.Msg)
		return
	}
	if adminRow.Password != req.Pwd {
		msg = retmsg.USER_PASWD_FAIL.Return()
		log.Warn(msg.Msg)
		return
	}
	now := time.Now().Unix()
	validSecond := int64(86400) // token有效期
	data := fmt.Sprintf("%v|%v", adminRow.Id, now+validSecond)
	token := auth.AesEncrypt(data)
	authInfo := &dstruct.Author{
		Username: req.Name,
		UserId:   adminRow.Id,
		Role:     adminRow.Role,
		Expired:  now + validSecond, // 24小时过期
		Token:    token,
	}
	session.InitSession(adminRow.Id, util.AssertMarshal(authInfo))
	msg.Data = authInfo
	updateMap := map[string]interface{}{
		"id":    authInfo.UserId,
		"token": token,
	}
	if err := common.UpdateTableById(db.GetAppDb(), &admin.TAdmin{}, updateMap); err != nil {
		log.Warn(err)
	}
	return
}
func logout(c *gin.Context) {
	msg := retmsg.OK.Return()
	defer func() {
		c.JSON(http.StatusOK, msg)
	}()

	authInfo, err := session.GetSession(c)
	if err != nil {
		log.Warn(err)
		return
	}
	updateMap := map[string]interface{}{
		"id":    authInfo.UserId,
		"token": "",
	}
	if err := common.UpdateTableById(db.GetAppDb(), &admin.TAdmin{}, updateMap); err != nil {
		log.Warn(err)
	}
	session.DeleteSession(authInfo.UserId)
	return
}

func changePwd(c *gin.Context) {
	msg := retmsg.OK.Return()
	defer func() {
		c.JSON(http.StatusOK, msg)
	}()

	authInfo, err := session.GetSession(c)
	if err != nil {
		log.Warn(err)
		msg = retmsg.USER_LOGOUT.Return()
		return
	}
	req := &struct {
		OldPwd string `json:"oldPwd"`
		NewPwd string `json:"newPwd"`
	}{}
	c.BindJSON(req)
	if len(req.OldPwd) != len(req.NewPwd) || len(req.OldPwd) != 32 {
		msg = retmsg.ERR_PARM.Return()
		log.Warn(msg.Msg)
		return
	}
	adminRow, err := admin.GetAdminById(authInfo.UserId)
	if err != nil {
		msg = retmsg.USER_NO_EXIST.Return()
		log.Warn(err)
		return
	}
	if adminRow.Password != req.OldPwd {
		msg = retmsg.USER_PASWD_FAIL.Return()
		log.Warn(err)
		return
	}
	updateMap := map[string]interface{}{
		"id":       authInfo.UserId,
		"password": req.NewPwd,
	}
	if err := common.UpdateTableById(db.GetAppDb(), &admin.TAdmin{}, updateMap); err != nil {
		log.Warn(err)
	}
	session.DeleteSession(authInfo.UserId)
	return
}
func addUser(c *gin.Context) {
	msg := retmsg.OK.Return()
	defer func() {
		c.JSON(http.StatusOK, msg)
	}()

	authInfo, err := session.GetSession(c)
	if err != nil {
		log.Warn(err)
		msg = retmsg.USER_LOGOUT.Return()
		return
	}
	_ = authInfo
	req := &struct {
		Name string `json:"name"`
		Pwd  string `json:"pwd"`
		Role string `json:"role"`
	}{}
	c.BindJSON(req)
	req.Name = strings.Trim(req.Name, " ")
	if len(req.Name) < 2 || len(req.Pwd) != 32 || (req.Role != consts.ROLE_SUPERADMIN &&
		req.Role != consts.ROLE_ADMIN && req.Role != consts.ROLE_APP && req.Role != consts.ROLE_GUEST) {
		msg = retmsg.ERR_PARM.Return()
		log.Warn(msg.Msg)
		return
	}
	if err := admin.AddAdmin(req.Name, req.Pwd, req.Role); err != nil {
		log.Warn(err)
		msg = retmsg.USER_ADD_FAIL.Return()
		return
	}
	msg.Data, _ = admin.SelectAdminList()
	return
}
func modifyUser(c *gin.Context) {
	msg := retmsg.OK.Return()
	defer func() {
		c.JSON(http.StatusOK, msg)
	}()
	authInfo, err := session.GetSession(c)
	if err != nil {
		log.Warn(err)
		msg = retmsg.USER_LOGOUT.Return()
		return
	}
	_ = authInfo
	req := &struct {
		UserId int    `json:"userId"`
		Name   string `json:"name"`
		Role   string `json:"role"`
		Status int    `json:"status"`
	}{}
	c.BindJSON(req)
	req.Name = strings.Trim(req.Name, " ")
	if req.UserId < 1 || len(req.Name) < 2 || (req.Role != consts.ROLE_SUPERADMIN &&
		req.Role != consts.ROLE_ADMIN && req.Role != consts.ROLE_APP && req.Role != consts.ROLE_GUEST) ||
		(req.Status != consts.USER_STATUS_ON && req.Status != consts.USER_STATUS_OFF) {
		msg = retmsg.ERR_PARM.Return()
		log.Warn(msg.Msg)
		return
	}
	if row, err := admin.SelectAdminByName(req.Name); err == nil && row.Id != req.UserId {
		msg = retmsg.USER_ALREADY.Return()
		log.Warn(msg.Msg)
		return
	}
	updateMap := map[string]interface{}{
		"id":       req.UserId,
		"username": req.Name,
		"role":     req.Role,
		"status":   req.Status,
	}
	if err := common.UpdateTableById(db.GetAppDb(), &admin.TAdmin{}, updateMap); err != nil {
		log.Warn(err)
		msg = retmsg.USER_ADD_FAIL.Return()
		return
	}
	msg.Data, _ = admin.SelectAdminList()
	return
}
func listUser(c *gin.Context) {
	msg := retmsg.OK.Return()
	defer func() {
		c.JSON(http.StatusOK, msg)
	}()

	authInfo, err := session.GetSession(c)
	if err != nil {
		log.Warn(err)
		msg = retmsg.USER_LOGOUT.Return()
		return
	}
	_ = authInfo
	rows, _ := admin.SelectAdminList()
	msg.Data = rows
	return
}
