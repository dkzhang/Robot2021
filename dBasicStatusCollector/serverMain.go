package main

import (
	"Robot2021/dBasicStatusCollector/collectWrite"
	pb "Robot2021/dBasicStatusCollector/grpc"
	"Robot2021/dBasicStatusCollector/server"
	myUtil "Robot2021/myUtils/formatTime"
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
