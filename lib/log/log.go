package log

import (
	"gin.go.tpl/lib/config"
	"gin.go.tpl/lib/constant"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	LogAPI  *Log
	LogOnce sync.Once
)

type Log struct {
	logger *logrus.Logger
	Conf   config.LogConf
}

func NewLogAPI(config config.LogConf) *Log {
	LogOnce.Do(func() {
		LogAPI = &Log{logrus.New(), config}
		LogAPI.SetLogrus()
	})
	return LogAPI
}

func (l *Log) GetLogger() *logrus.Logger {
	return l.logger
}

func (l Log) SetLogrus() {
	if l.Conf.Level < 7 {
		l.logger.SetLevel(logrus.Level(l.Conf.Level))
	}
	l.logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: constant.DefaultTimestampFormat, DisableColors: false,
		ForceColors: true, FullTimestamp: true})
	if l.Conf.LogFile != "" { // 如果日志文件非空, 将日志打到对应文件
		file, err := os.OpenFile(l.Conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			// 如果输出到文件, 将格式设置为json
			l.logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: constant.DefaultTimestampFormat})
			l.logger.SetOutput(file)
		} else {
			l.Info("记录文件异常, 采用默认标准错误输出", err)
		}
	}
}

func (l Log) Print(args ...interface{}) {
	l.logger.Println(args)
}

func (l Log) Trace(args ...interface{}) {
	l.logger.Traceln(args)
}

func (l Log) Debug(args ...interface{}) {
	l.logger.Debugln(args)
}

func (l Log) Info(args ...interface{}) {
	l.logger.Infoln(args)
}

func (l Log) Warn(args ...interface{}) {
	l.logger.Warnln(args)
}

func (l Log) Error(args ...interface{}) {
	l.logger.Errorln(args)
}

func (l Log) Fatal(args ...interface{}) {
	l.logger.Fatalln(args)
}

func (l Log) Panic(args ...interface{}) {
	logrus.Panicln(args)
}
