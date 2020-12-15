package server

import (
	_ "encoding/json"
	_ "github.com/gomodule/redigo/redis"
	"testing"
)

func TestServer_GetRobotStatus(t *testing.T) {

	server := BasicStatusServer{}

	reply, err := server.GetRobotStatus(nil, nil)
	if err != nil {
		t.Errorf("GetRobotStatus error: %v", err)
	} else {
		t.Logf("GetRobotStatus reply = %v", *reply)
	}
}
