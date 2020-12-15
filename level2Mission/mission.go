package level2Mission

import "time"

type Mission struct {
	Type       int
	Name       string
	Parameters string
	CreatedAt  time.Time
}

const (
	EndMission = -1
)
