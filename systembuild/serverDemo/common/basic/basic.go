package basic

import (
	"fmt"
	"os"
	"path/filepath"
	"serverDemo/common/log"
	"strconv"

	"github.com/widuu/goini"
)

// MysqlType 存储数据库相关配置
type MysqlType struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type RedisType struct {
	Host     string
	Password string
	Port     string
}

type AppType struct {
	Port    string
	RpcPort string
	Secret  string
	RunCron bool
	BinName string
}
type PathType struct {
	LogDir string
}
type Logger struct {
	Filename string `json:"filename"`
	Level    int    `json:"level"`
	Maxlines int    `json:"maxlines"`
	Maxsize  int    `json:"maxsize"`
	Daily    bool   `json:"daily"`
	Maxdays  int    `json:"maxdays"`
	Color    bool   `json:"color"`
}

var (
	MysqlApp     MysqlType
	MysqlMonitor MysqlType
	App          AppType
	Path         PathType
	Redis        RedisType
)

func init() {
	initIni()
	initDir()
	initLog()
}
func initDir() {
	os.MkdirAll(Path.LogDir, os.ModePerm)
}
func initLog() {
	fileName := filepath.Join(Path.LogDir, "console.log")
	logConfig := log.LogConfig{
		Filename:        fileName,
		RetainFileCount: 2048,
	}
	log.SetLogger(logConfig)
	log.SetLevel(log.DebugLevel)
}

func initIni() {
	conf := goini.SetConfig(filepath.Join("config", "conf.ini"))
	MysqlApp.User = conf.GetValue("mysqlApp", "user")
	MysqlApp.Password = conf.GetValue("mysqlApp", "password")
	MysqlApp.Host = conf.GetValue("mysqlApp", "host")
	MysqlApp.Port = conf.GetValue("mysqlApp", "port")
	MysqlApp.Database = conf.GetValue("mysqlApp", "database")

	MysqlMonitor.User = conf.GetValue("mysqlMonitor", "user")
	MysqlMonitor.Password = conf.GetValue("mysqlMonitor", "password")
	MysqlMonitor.Host = conf.GetValue("mysqlMonitor", "host")
	MysqlMonitor.Port = conf.GetValue("mysqlMonitor", "port")
	MysqlMonitor.Database = conf.GetValue("mysqlMonitor", "database")

	Redis.Password = conf.GetValue("redis", "password")
	Redis.Host = conf.GetValue("redis", "host")
	Redis.Port = conf.GetValue("redis", "port")

	App.Port = conf.GetValue("app", "port")
	App.RpcPort = conf.GetValue("app", "rpcPort")
	App.Secret = conf.GetValue("app", "secret")
	App.BinName = conf.GetValue("app", "binName")
	var err error
	if App.RunCron, err = strconv.ParseBool(conf.GetValue("app", "runcron")); err != nil {
		App.RunCron = true
	}

	Path.LogDir, _ = filepath.Abs(conf.GetValue("path", "logsdir"))

	checkIni()
}

func checkIni() {
	if len(Path.LogDir) < 3 {
		err := fmt.Errorf("logsdir err in conf.ini")
		fmt.Println(err)
		panic(err)
	}
}
