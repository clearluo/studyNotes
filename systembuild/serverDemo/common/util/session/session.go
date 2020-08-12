package session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"serverDemo/common/auth"
	"serverDemo/common/cache"
	"serverDemo/common/dstruct"
	"serverDemo/common/log"
	"serverDemo/common/util"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func InitSession(user_id int, datastr string) {
	log.Info(user_id, " init_session", datastr)
	cache.Cache.Set("@login_"+strconv.Itoa(user_id), datastr, 240*time.Hour)
}
func DeleteSession(userId int) {
	cache.Cache.Delete("@login_" + strconv.Itoa(userId))
}
func GetSession(c *gin.Context) (*dstruct.Author, error) {
	data, err := c.GetRawData()
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	now := time.Now().Unix()
	authInfo := &dstruct.Author{}
	userIdStr := c.GetHeader("userId")
	token := c.GetHeader("token")
	log.Infof("reqUrl:%v userId:%v token:%v body:%v\n", c.Request.URL, userIdStr, token, string(data))
	if len(token) < 5 {
		err := fmt.Errorf("token is too short:%v", token)
		log.Warn(err)
		return nil, err
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	jsonstr, flag := cache.Cache.Get("@login_" + userIdStr)
	if !flag || jsonstr == nil {
		log.Info("缓存无数据，查数据库token")
		//rowAdmin, err := admin.GetAdminById(userId)
		//if err != nil || len(rowAdmin.Token) < 10 {
		//	err := fmt.Errorf("token not found")
		//	log.Warn(err)
		//	return nil, err
		//}
		//authInfo.UserId = userId
		//authInfo.Username = rowAdmin.Username
		//authInfo.Role = rowAdmin.Role
		//authInfo.Token = rowAdmin.Token
		//doUpdate("Admin", map[string]interface{}{"token": tokenDb}, map[string]interface{}{"id": user_id})
		InitSession(userId, util.AssertMarshal(authInfo))
	} else {
		json.Unmarshal([]byte(jsonstr.(string)), &authInfo)
	}
	log.Infof("%v session=%v", userId, authInfo)
	if authInfo.Token != token {
		err := fmt.Errorf("auth.Token!=token")
		log.Warn(err)
		//return nil, err // TODO 测试先不校验

	}
	// 验证token是否过期
	tokenData := auth.AesDecrypt(token)
	tokenSli := strings.Split(tokenData, "|")
	if len(tokenSli) != 2 {
		err := fmt.Errorf("token format err:%v", tokenSli)
		log.Warn(err)
		return nil, err
	}
	expired, err := strconv.ParseInt(tokenSli[1], 10, 64)
	if err != nil {
		log.Warn(err, tokenSli[1])
		return nil, err
	}
	if expired < now {
		err := fmt.Errorf("token expired:%v,%v", expired, now)
		log.Warn(err)
		// return nil, err // TODO 测试先不校验
	}
	return authInfo, nil
}
