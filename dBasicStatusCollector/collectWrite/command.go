package collectWrite

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateSubscribeRobotStatusCommand(pFrequency *float64) (name, cmd string, uuid string) {
	name = "/api/request_data"
	cmd = "/api/request_data?topic=robot_status"

	if pFrequency != nil {
		cmd += fmt.Sprintf("&frequency=%f", *pFrequency)
	}

	rand.Seed(time.Now().Unix())
	uuid = fmt.Sprintf("%X", rand.Uint32())
	cmd += fmt.Sprintf("&uuid=%s", uuid)

	return name, cmd, uuid
}
