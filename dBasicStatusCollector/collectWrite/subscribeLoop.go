package collectWrite

import (
	"Robot2021/chassisDriver/common"
	"Robot2021/chassisDriver/socketCommunication"
	"Robot2021/databaseCommon/redisOps"
	myUtil "Robot2021/myUtils/formatTime"
	"Robot2021/myUtils/splitJSON"
	"fmt"
	"log"
	"time"
)

func SubscribeLoop() {

	//实例化一个通信模块
	serverIPandPort := "192.168.10.10:31001"
	psm := socketCommunication.SocketManagementFactory(serverIPandPort)
	defer psm.Cancel()

	//构造定期获取机器人状态命令（附随机数）并发送
	cmdSubscribeStruct := socketCommunication.CommandStruct{Name: "Subscribe Robot Status"}
	cmdSubscribeStruct.Name, cmdSubscribeStruct.Command, cmdSubscribeStruct.UUID = GenerateSubscribeRobotStatusCommand(nil)
	psm.CommandChan <- &cmdSubscribeStruct

	//循环接收传回的消息
	cmdSubscribeFlag := false
	var err error

	opts := &redisOps.RedisOpts{
		Host: redisOps.RedisHost,
	}
	theRedis := redisOps.NewRedis(opts)

	for {
		select {
		case strResult := <-psm.ResultChan:

			for _, strJSON := range splitJSON.Split(strResult) {

				if cmdSubscribeFlag == false {
					//命令解析
					//检查是否为所发命令的回复
					cmdSubscribeFlag, err = CmdResponseParse(strJSON, cmdSubscribeStruct.Name, cmdSubscribeStruct.UUID)

					if err != nil {
						log.Printf("CmdResponseParse error: %v", err)
					} else if cmdSubscribeFlag == true {
						continue
					}
				}

				//订阅消息解析
				//检查是否为订阅消息robot status，如果是，则存入redis
				log.Printf("check for subscribe: %s", strJSON)

				if err = SubscribeResponseParse(theRedis, strJSON); err != nil {
					log.Printf("SubscribeResponseParse error: %v", err)
				}

			}

		case feedback := <-psm.FeedbackChan:
			log.Printf("socketCommunication feedback: %v", feedback)
		}
	}

	return
}

// 解析命令的响应报文
// 如果收到的不是所发命令对应的响应报文，且没有出错，返回false，nil
// 如果是对应的响应报文，且状态正确，返回true，nil
// 如果出错，返回error
func CmdResponseParse(result string, name, uuid string) (bool, error) {
	// 检查是否为所发命令的回复
	pcr, err := common.CommandDetection(result, name, uuid)
	//检查是否出错
	if err != nil {
		return false, fmt.Errorf("CommandDetection error: %v", err)
	} else {
		if pcr != nil {
			if pcr.Status == "OK" {
				log.Printf("Command OK: %v", pcr)
				return true, nil
			} else {
				return false, fmt.Errorf("respone error: %v", pcr)
			}
		}
	}
	return false, nil
}

// 解析命令的响应报文
// 如果收到的不是callback消息订阅，或者不是robot_status主题，或者运动没有完成，返回false
// 如果是对应的响应报文，且指示移动命令已完成，返回true
// 如果出错，由于订阅消息的可重复性，直接忽略，返回false
func SubscribeResponseParse(theRedis *redisOps.Redis, result string) error {

	pct, err := common.CallbackTopicDetection(result, "robot_status")

	//检查是否出错
	if err != nil {
		return fmt.Errorf("CallbackTopicDetection error: %v", err)
	} else {
		if pct != nil {
			//是robot status的订阅消息,则进一步解析
			robotStatus := RobotStatusTopic{}
			err = robotStatus.UnmarshalJSON(result)
			if err != nil {
				return fmt.Errorf("robotStatus.UnmarshalJSON error: %v", err)
			} else {
				robotStatus.TimeStamp = myUtil.FormatTime(time.Now())
				//write to redis
				theRedis.ListMaxLenRPush("RobotStatus", robotStatus, 256)
			}
		}
	}
	return nil
}
