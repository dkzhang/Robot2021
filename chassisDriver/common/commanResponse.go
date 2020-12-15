package common

import (
	"encoding/json"
	"fmt"
)

type ResultType struct {
	Type string `json:"type"`
}

type CommandResponse struct {
	Type         string `json:"type"`
	Command      string `json:"Command"`
	ErrorMessage string `json:"error_message"`
	Status       string `json:"status"`
	UUID         string `json:"uuid"`
}

//type CommandResult struct {
//	BasicInfo BasicCommandResult
//	StrJSON   string
//}

func CommandDetection(strJSON string, name, uuid string) (pcr *CommandResponse, err error) {
	//先解析消息类型
	rt := ResultType{}
	err = json.Unmarshal([]byte(strJSON), &rt)
	if err != nil {
		return nil, fmt.Errorf("result type json unmarshal error, strJSON = %s, err = %v", strJSON, err)
	}

	//判断消息类型是否为response
	if rt.Type != "response" {
		return nil, nil
	}

	//按response格式进行详细解析
	cr := CommandResponse{}
	err = json.Unmarshal([]byte(strJSON), &cr)
	if err != nil {
		return nil, fmt.Errorf("CommandResponse json unmarshal error: %v", err)
	}

	//判断是否所发命令
	if cr.Command != name {
		return nil, nil
	}

	//判断是否给定uuid。如果给定，判断uuid是否相符
	if len(uuid) != 0 && cr.UUID != uuid {
		return nil, nil
	}

	return &cr, nil
}
