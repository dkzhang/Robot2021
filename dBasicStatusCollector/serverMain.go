package main

import (
	pb "Robot2019/dataServer/robotStatusServer/grpc"
	"Robot2019/myUtil"
	"Robot2021/dBasicStatusCollector/collectWrite"
	"Robot2021/dBasicStatusCollector/server"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	// launch the collect & write loop
	go collectWrite.SubscribeLoop()

	// launch the grpc server
	const port = ":50071"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf(" fatal error! failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRobotStatusServiceServer(s, &server.BasicStatusServer{})
	fmt.Printf("Begin to serve %s", myUtil.FormatTime(time.Now()))
	if err := s.Serve(lis); err != nil {
		log.Printf(" fatal error! failed to serve: %v", err)
	}

}
