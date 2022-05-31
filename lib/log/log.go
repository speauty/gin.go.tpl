package log

import (
	"gin.go.tpl/lib/config"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	LogAPI  *Log
	LogOnce sync.Once
)

type Log struct{}

func NewLogAPI(config config.LogConf) *Log {
	LogOnce.Do(func() {
		LogAPI = &Log{}
		if config.Level < 7 {
			LogAPI.SetLevel(logrus.Level(config.Level))
		}
	})
	return LogAPI
}

func (l Log) SetLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

func (l Log) Print(args ...interface{}) {
	logrus.Println(args)
}

func (l Log) Trace(args ...interface{}) {
	logrus.Traceln(args)
}

func (l Log) Debug(args ...interface{}) {
	logrus.Debugln(args)
}

func (l Log) Info(args ...interface{}) {
	logrus.Infoln(args)
}

func (l Log) Warn(args ...interface{}) {
	logrus.Warnln(args)
}

func (l Log) Error(args ...interface{}) {
	logrus.Errorln(args)
}

func (l Log) Fatal(args ...interface{}) {
	logrus.Fatalln(args)
}

func (l Log) Panic(args ...interface{}) {
	logrus.Panicln(args)
}
