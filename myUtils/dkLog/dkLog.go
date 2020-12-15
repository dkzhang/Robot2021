package dkLog

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var dlf *logrus.Logger
var once sync.Once

func LogF() *logrus.Logger {
	once.Do(func() {
		dlf = loggerToFile("", "robot2021")
	})
	return dlf
}
