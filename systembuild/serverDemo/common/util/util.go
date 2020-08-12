package util

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"serverDemo/common/auth"
	"serverDemo/common/dstruct"
	"serverDemo/common/log"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
)

var localzone *time.Location

func init() {
	localzone, _ = time.LoadLocation("Asia/Shanghai")
}

func AssertMarshal(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// 获取整点以来的秒数
func GetSecondFromHour() int64 {
	return time.Now().Unix() % 3600
}

//获取整天以来的秒数
func GetSecondFromDay() int64 {
	return time.Now().Unix() - GetSecondByDay00()
}

//获取整周以来的秒数
func GetSecondFromWeek() int64 {
	day := int64(time.Now().Weekday())
	return 86400*(day-1) + GetSecondFromDay()
}

// 获取整月以来的秒数
func GetSecondFromMonth() int64 {
	day := int64(time.Now().Day())
	return 86400*(day-1) + GetSecondFromDay()
}

//获取当前日期(20170802)零点对应的Unix时间戳
func GetSecondByDay00() int64 {
	timeStr := time.Now().Format("2006-01-02")
	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.Unix()
}

// 获取当前日期前后n天对应的日期证书,0代表获取当前日期整数
func GetDateByN(n int) int64 {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, n)
	dayStr := yesTime.Format("20060102")
	day, _ := strconv.ParseInt(dayStr, 10, 64)
	return day
}

// 根据时间戳获取对应日期整数
func GetTDayByUnixTime(nowUnix int64) int64 {
	if nowUnix < 1 {
		return 0
	}
	tm := time.Unix(nowUnix, 0)
	nowDay, err := strconv.ParseInt(tm.Format("20060102"), 10, 64)
	if err != nil {
		log.Error(err)
		return 0
	}
	return nowDay
}

func ParseToken(tokenStr string) (*dstruct.TokenDsta, error) {
	retErr := errors.New("parse token err")
	if len(tokenStr) < 1 {
		return nil, retErr
	}
	srcToken := auth.AesDecrypt(tokenStr)
	if len(srcToken) < 10 {
		return nil, retErr
	}
	srcSplic := strings.Split(srcToken, "|")
	if len(srcSplic) != 3 && len(srcSplic) != 4 {
		return nil, retErr
	}
	userId, err := strconv.ParseInt(srcSplic[0], 10, 64)
	if err != nil || userId < 1 {
		log.Warn(err)
		return nil, retErr
	}
	milliSecond, err := strconv.ParseInt(srcSplic[1], 10, 64)
	if err != nil || milliSecond < 1 {
		return nil, retErr
	}
	areaId, err := strconv.ParseInt(srcSplic[2], 10, 64)
	if err != nil {
		log.Warn(err)
		return nil, retErr
	}
	tokenData := &dstruct.TokenDsta{
		UserId:      int(userId),
		MilliSecond: milliSecond,
		AreaId:      int(areaId),
	}
	if len(srcSplic) == 4 {
		tokenData.AuthFlg, _ = strconv.Atoi(srcSplic[3])
	}
	curtime := time.Now().Unix()
	if curtime-tokenData.MilliSecond/1000 > 600 {
		log.Warn("token失效:", curtime, tokenData.MilliSecond/1000)
		//return nil, retErr
	}
	return tokenData, nil
}

// 统计某函数执行时间
// 使用方式
// defer utils.Profiling("test")()
func Profiling(msg string) func() {
	start := time.Now()
	return func() {
		log.Infof(fmt.Sprintf("%s[%s]:%s", msg, "use", time.Since(start)))
	}
}

func PostForm(c *gin.Context, mykey string, def string) string {
	myvalue := c.PostForm(mykey)
	if myvalue == "" {
		return def
	} else {
		return myvalue
	}
}

func PostFormInt(c *gin.Context, mykey string, def int) int {
	myvalue := c.PostForm(mykey)
	if myvalue == "" {
		myvalue = c.Query(mykey)
	}
	if myvalue == "" {
		return def
	} else {
		res, _ := strconv.Atoi(myvalue)
		return res
	}
}

// 判断文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetFieldName(t reflect.Type) (map[string]string, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		err := fmt.Errorf("Check type error not Struct")
		fmt.Println(err)
		log.Warn(err)
		return nil, err
	}
	fieldNum := t.NumField()
	result := make(map[string]string, 0)
	for i := 0; i < fieldNum; i++ {
		result[t.Field(i).Tag.Get("json")] = t.Field(i).Type.Name()
	}
	return result, nil
}

type TplJson struct {
	Length int                      `json:"length"`
	Data   []map[string]interface{} `json:"data"`
}

// url encode string, is + not %20
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

// url decode string
func UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

// base64 encode
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// base64 decode
func Base64Decode(str string) (string, error) {
	s, e := base64.StdEncoding.DecodeString(str)
	return string(s), e
}

// s:结构体
// columnType:结构体中的成员
// columnValue:成员对应的值
// 此函数将指定的结构体成员值更新到结构体中
func SetStructValueByType(s interface{}, columnType string, columnValue interface{}) error {
	columnValueV := reflect.ValueOf(columnValue)
	var setValue reflect.Value
	var flag bool = false
	//t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	//length := t.Elem().NumField()
	for i, n := 0, v.Elem().NumField(); i < n; i++ {
		//if t.Elem().Field(i).Name == columnType {
		if v.Elem().Type().Field(i).Name == columnType {
			setValue = v.Elem().Field(i)
			flag = true
			break
		}
	}
	if !flag {
		return errors.New("struct is not type:")
	} else if !setValue.CanSet() {
		return errors.New("setValue.CanSet is false")
	} else if setValue.Kind() != columnValueV.Kind() {
		return errors.New("struct field and value of type is error")
	}
	switch columnValueV.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		setValue.SetInt(int64(columnValueV.Int()))
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		setValue.SetUint(uint64(columnValueV.Uint()))
	case reflect.Float32, reflect.Float64:
		setValue.SetFloat(float64(columnValueV.Float()))
	case reflect.String:
		setValue.SetString(columnValueV.String())
	default:
		return errors.New("columnValue err for:" + columnType)
	}
	return nil
}

func GetDateFormat1() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02")
}

func GetRand(num int) int {
	if num < 1 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(num)
}

func GetHtml(url string) ([]byte, error) {
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Warn("get error", url, err.Error())
		return []byte{}, err
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	return bodyBytes, nil
}

func EncodeMd5(seed string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(seed))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func IsValiUser(username string) bool {
	ok, err := regexp.MatchString("^[a-zA-Z0-9_@]+$", username)
	if err != nil {
		log.Warn(err)
		return false
	}
	return ok
}

func LogErr(err error) {
	if err != nil {
		log.Warn("err", err.Error())
	}
}

func ValidBetweenDate(startDate string, endDate string) error {
	startTime, err := time.Parse("2006-01-02 15:04:05", startDate)
	if err != nil {
		return err
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", endDate)
	if err != nil {
		return err
	}
	startDay, endDay := startTime.Unix()/86400, endTime.Unix()/86400
	if day := endDay - startDay; day > 365 || day < 0 {
		return fmt.Errorf("startDate and endDate between err")
	}
	return nil
}

func GetTimeByYyyymmddhhmm() int {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	minuteStr := tm.Format("200601021504")
	minuteInt, _ := strconv.ParseInt(minuteStr, 10, 64)
	return int(minuteInt)
}

func ParseTime(str string) (time.Time, error) {
	return time.ParseInLocation("200601021504", fmt.Sprintf("%v", str), localzone)
}

func ParseTime2(str string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%v", str), localzone)
}

func ExecBash(str string) (string, error) {
	log.Info("ExecBash:", str)
	str = DeleteExtraSpace(str)
	var out []byte
	var err error
	cmdSli := strings.Split(str, "|")
	for _, cmdStr := range cmdSli {
		cmdStr = strings.Trim(cmdStr, " ")
		param := strings.Split(cmdStr, " ")
		if len(param) < 1 {
			continue
		}
		for i := range param {
			param[i] = strings.ReplaceAll(param[i], "&n"+
				"bsp", " ")
		}
		cmd := exec.Command(param[0], param[1:]...)
		cmd.Stdin = bytes.NewBuffer(out)
		out, err = cmd.CombinedOutput()
		if err != nil {
			break
		}
	}
	log.Info("BashResult:", string(out))
	return string(out), err
}

func DeleteExtraSpace(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "\t", " ", -1) //替换tab为空格
	s1 = strings.Trim(s1, " ")
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}

func GrepWin(searchContent, filePath string) (string, error) {
	if runtime.GOOS == "windows" {
		command := exec.Command("cmd", "/C", "findstr /i /c:"+searchContent+" "+filePath)
		bytes, e := command.Output()
		if e != nil {
			return "", e
		}
		result := (*string)(unsafe.Pointer(&bytes))
		return *result, nil
	} else {
		command := exec.Command("sh", "-c", "grep "+searchContent+" "+filePath)
		bytes, e := command.Output()
		if e != nil {
			return "", e
		}
		result := (*string)(unsafe.Pointer(&bytes))
		return *result, nil
	}
}

func TailWin(filename string, n int) (ret string, err error) {
	lines := []string{}
	const (
		defaultBufSize = 4096
	)
	f, e := os.Stat(filename)
	if e == nil {
		size := f.Size()
		var fi *os.File
		fi, err = os.Open(filename)
		if err == nil {
			b := make([]byte, defaultBufSize)
			sz := int64(defaultBufSize)
			nn := n
			bTail := bytes.NewBuffer([]byte{})
			istart := size
			flag := true
			for flag {
				if istart < defaultBufSize {
					sz = istart
					istart = 0
				} else {
					istart -= sz
				}
				_, err = fi.Seek(istart, os.SEEK_SET)
				if err == nil {
					mm, e := fi.Read(b)
					if e == nil && mm > 0 {
						j := mm
						for i := mm - 1; i >= 0; i-- {
							if b[i] == '\n' {
								bLine := bytes.NewBuffer([]byte{})
								//bLine.Write( b[i+1:j] )
								bLine.Write(b[i:j])
								j = i
								if bTail.Len() > 0 {
									bLine.Write(bTail.Bytes())
									bTail.Reset()
								}

								if (nn == n && bLine.Len() > 0) || nn < n { //skip last "\n"
									lines = append(lines, bLine.String())
									nn--
								}
								if nn == 0 {
									flag = false
									break
								}
							}
						}
						if flag && j > 0 {
							if istart == 0 {
								bLine := bytes.NewBuffer([]byte{})
								bLine.Write(b[:j])
								if bTail.Len() > 0 {
									bLine.Write(bTail.Bytes())
									bTail.Reset()
								}
								lines = append(lines, bLine.String())
								flag = false
							} else {
								bb := make([]byte, bTail.Len())
								copy(bb, bTail.Bytes())
								bTail.Reset()
								bTail.Write(b[:j])
								bTail.Write(bb)
							}
						}
					}
				}
			}
		}
		defer fi.Close()
	}
	ret = strings.Join(lines, "\n")
	ret = strings.Trim(ret, "\n")
	return
}

// 获取正在运行的上级函数名
func RunFatherFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
