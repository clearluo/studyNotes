package hello

import (
	"context"
	"serverDemo/common/log"
	"serverDemo/common/myredis"
	pb "serverDemo/common/proto/helloworld"
	"serverDemo/common/util"
)

type HelloServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	myredis.GetDb().Incr(myredis.GetMonitorQpsKey(util.GetTimeByYyyymmddhhmm()))
	log.Infof("Received: %v\n", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
