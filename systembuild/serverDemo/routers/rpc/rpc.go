package rpc

import (
	"fmt"
	"net"
	"serverDemo/common/basic"
	"serverDemo/common/log"
	"serverDemo/common/myredis"
	"serverDemo/common/proto/helloworld"
	monitorPb "serverDemo/common/proto/monitor"
	"serverDemo/common/util"
	"serverDemo/handler/rpc/hello"
	"serverDemo/handler/rpc/monitor"

	"google.golang.org/grpc"
)

func InitRpcServer() {
	myredis.GetDb().Incr(myredis.GetMonitorQpsKey(util.GetTimeByYyyymmddhhmm()))
	defer func() {
		if err := recover(); err != nil {
			log.Error("RpcServer.panic:", err)
		}
	}()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", basic.App.RpcPort))
	if err != nil {
		err := fmt.Errorf("failed to listen:%v", err)
		log.Warn(err)
		fmt.Println(err)
		panic(err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &hello.HelloServer{})
	monitorPb.RegisterMonitorServer(s, &monitor.MonitorServer{})
	log.Infof("start rpc server @port=%v\n", basic.App.RpcPort)
	if err := s.Serve(lis); err != nil {
		err := fmt.Errorf("failed to serve:%v", err)
		log.Warn(err)
		fmt.Println(err)
		panic(err)
	}
}
