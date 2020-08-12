package main

// import "C"
import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"serverDemo/common/basic"
	"serverDemo/common/basic/runbefore"
	_ "serverDemo/common/cache"
	"serverDemo/common/log"
	_ "serverDemo/common/myredis"
	_ "serverDemo/db"
	"serverDemo/routers/http"
	"serverDemo/routers/rpc"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli"
)

var pidPath string

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	app := cli.NewApp()

	pidPath = basic.App.BinName + ".pid"

	app.Name = "hitball"
	app.Author = "ndc authors"
	app.Version = "0.0.1"
	app.Copyright = "ndc authors reserved"
	app.Usage = "hitball start|stop|restart|{-d}"

	app.Commands = []cli.Command{
		cli.Command{
			Name: "start",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "d", Usage: "run background"},
			},
			Aliases: []string{"start"},
			Action:  start,
		},
		cli.Command{
			Name:    "stop",
			Aliases: []string{"stop"},
			Action:  stop,
		},
		cli.Command{
			Name:    "restart",
			Aliases: []string{"restart"},
			Action:  restart,
		},
	}
	app.Run(os.Args)
}

func start(ctx *cli.Context) error {
	d := ctx.Bool("d")
	log.Info("start_d ", d)
	if d {
		start := exec.Command(basic.App.BinName, "start")
		start.Start()
		os.Exit(0)
	}
	doStart()
	return nil
}

// Start 启动
func doStart() {
	pid := fmt.Sprintf("%d", os.Getpid())
	if err := ioutil.WriteFile(pidPath, []byte(pid), 0666); err != nil {
		log.Warn("start pid error ", pid)
		panic(err)
	}
	runbefore.InitRun()
	fmt.Println("RunCron:", basic.App.RunCron)
	go rpc.InitRpcServer()
	http.InitRoute()
}

// Stop 停止
func stop(ctx *cli.Context) error {
	pid, _ := ioutil.ReadFile(pidPath)
	log.Info("doStop ", ctx.String("stop"), string(pid))
	cmd := exec.Command("kill", "-9", string(pid))
	cmd.Start()
	if err := ioutil.WriteFile(pidPath, nil, 0666); err != nil {
		log.Warnf("write pid error %v %v", string(pid), err.Error())
	}
	return nil
}

// Restart 重启
func restart(ctx *cli.Context) error {
	log.Info("doRestart ", ctx.String("restart"))
	pid, _ := ioutil.ReadFile(pidPath)
	log.Info("restarting..." + string(pid))
	stop := exec.Command("kill", "-9", string(pid))
	stop.Start()
	start := exec.Command(basic.App.BinName, "start")
	start.Start()
	return nil
}
