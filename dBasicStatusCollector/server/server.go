package server

import (
	"Robot2019/cache"
	robotStatusWriter "Robot2019/chassisDriverForRobot/subscribeRobotStatusWriter/server"
	pb "Robot2019/dataServer/robotStatusServer/grpc"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type BasicStatusServer struct {
	pb.UnimplementedRobotStatusServiceServer
}

func (s *BasicStatusServer) GetRobotStatus(ctx context.Context, in *pb.RobotStatusRequest) (*pb.RobotStatusReply, error) {
	log.Printf("Received: %v", in.GetTag())

	opts := &cache.RedisOpts{
		Host: cache.RedisHost,
	}
	theRedis := cache.NewRedis(opts)

	result, err := theRedis.ListIndex("RobotStatus", -1)
	if err != nil {
		return nil, fmt.Errorf("theRedis.ListIndex: %v", err)
	}

	robotStatus := robotStatusWriter.RobotStatusTopic{}
	err = json.Unmarshal([]byte(result), &robotStatus)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %v", err)
	} else {
		return &pb.RobotStatusReply{
			MoveStatus:      robotStatus.Results.MoveStatus,
			ChargeStatus:    robotStatus.Results.ChargeState,
			SoftEstopStatus: robotStatus.Results.SoftEStopState,
			HardEstopStatus: robotStatus.Results.HardEStopState,
			PowerPercent:    int64(robotStatus.Results.PowerPercent),
			X:               robotStatus.Results.CurrentPose.X,
			Y:               robotStatus.Results.CurrentPose.Y,
			Theta:           robotStatus.Results.CurrentPose.Theta,
			Datetime:        robotStatus.TimeStamp,
			ErrorMessage:    "",
		}, nil
	}
}
