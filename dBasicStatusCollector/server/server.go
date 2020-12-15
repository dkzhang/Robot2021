package server

import (
	"Robot2021/dBasicStatusCollector/collectWrite"
	pb "Robot2021/dBasicStatusCollector/grpc"
	"Robot2021/databaseCommon/redisOps"
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

	opts := &redisOps.RedisOpts{
		Host: redisOps.RedisHost,
	}
	theRedis := redisOps.NewRedis(opts)

	result, err := theRedis.ListIndex("RobotStatus", -1)
	if err != nil {
		return nil, fmt.Errorf("theRedis.ListIndex: %v", err)
	}

	robotStatus := collectWrite.RobotStatusTopic{}
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
