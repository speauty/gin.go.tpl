package log

import (
	"gin.go.tpl/lib/config"
	"gin.go.tpl/lib/constant"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	LogApi  *Log
	LogOnce sync.Once
)

type Log struct {
	logger *logrus.Logger
	Conf   config.LogConf
}

func NewLogApi(config config.LogConf) *Log {
	LogOnce.Do(func() {
		LogApi = &Log{logrus.New(), config}
		LogApi.SetLogrus()
	})
	return LogApi
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
		var fd *rotatelogs.RotateLogs
		if l.Conf.LogRotationCount > 0 {
			fd, _ = rotatelogs.New(
				l.Conf.LogFile+".%Y%m%d",
				rotatelogs.WithLinkName(l.Conf.LogFile),
				rotatelogs.WithMaxAge(time.Duration(l.Conf.LogMaxAge)*time.Second),
				rotatelogs.WithRotationCount(l.Conf.LogRotationCount),
			)
		} else {
			if l.Conf.LogRotationTime == 0 {
				l.Conf.LogRotationTime = 60 * 60 * 24
			}
			fd, _ = rotatelogs.New(
				l.Conf.LogFile+".%Y%m%d",
				rotatelogs.WithLinkName(l.Conf.LogFile),
				rotatelogs.WithMaxAge(time.Duration(l.Conf.LogMaxAge)*time.Second),
				rotatelogs.WithRotationTime(time.Duration(l.Conf.LogRotationTime)*time.Second),
			)
		}

		l.logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: constant.DefaultTimestampFormat})
		l.logger.SetOutput(fd)
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
