package level2Mission

import "time"

func GenMission_End() Mission {
	return Mission{
		Type:       EndMission,
		Name:       "结束2级任务",
		Parameters: "",
		CreatedAt:  time.Now(),
	}
}
