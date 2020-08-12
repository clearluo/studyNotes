package monitor

import (
	"context"
	"serverDemo/common/log"
	pb "serverDemo/common/proto/monitor"
)

type MonitorServer struct {
	pb.UnimplementedMonitorServer
}

func (s *MonitorServer) Report(ctx context.Context, in *pb.ReportRequest) (*pb.ReportRequestReply, error) {
	log.Infof("Received: %v %v\n", in.GetAction(), in.GetNode())
	if len(in.GetAction()) < 2 || len(in.GetNode()) < 2 {
		return &pb.ReportRequestReply{
			Code: 1022,
			Msg:  "参数错误",
		}, nil
	}

	return &pb.ReportRequestReply{
		Code: 200,
		Msg:  "success",
	}, nil
}

func (s *MonitorServer) ReportBak(ctx context.Context, in *pb.ReportRequest) (*pb.ReportRequestReply, error) {
	log.Infof("Received: %v\n", in.String())
	if len(in.GetAction()) < 2 || len(in.GetNode()) < 2 {
		return &pb.ReportRequestReply{
			Code: 1022,
			Msg:  "参数错误",
		}, nil
	}
	return &pb.ReportRequestReply{
		Code: 200,
		Msg:  "success",
	}, nil
}
