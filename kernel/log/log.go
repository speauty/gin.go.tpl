package log

import (
	"gin.go.tpl/kernel/cfg"
	"gin.go.tpl/kernel/constant"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"runtime"
	"sync"
	"time"
)

var (
	api  *Log
	once sync.Once
)

type Log struct {
	logger *logrus.Logger
	cfg    *cfg.LogConf
}

func NewLogApi(cfg *cfg.LogConf) *Log {
	once.Do(func() {
		api = &Log{logrus.New(), cfg}
		api.SetLogrus()
	})
	return api
}

func (l *Log) GetLogger() *logrus.Logger {
	return l.logger
}

func (l Log) SetLogrus() {
	if l.cfg.Level < 7 {
		l.logger.SetLevel(logrus.Level(l.cfg.Level))
	}
	l.logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: constant.DefaultTimestampFormat, DisableColors: false,
		ForceColors: true, FullTimestamp: true})
	if l.cfg.LogFile != "" { // 如果日志文件非空, 将日志打到对应文件
		var fd *rotatelogs.RotateLogs
		optLogFileFmt := l.cfg.LogFile + ".%Y%m%d"
		optWithLinkName := rotatelogs.WithLinkName(l.cfg.LogFile)
		optWithMaxAge := rotatelogs.WithMaxAge(time.Duration(l.cfg.LogMaxAge) * time.Second)
		optWithRotationCount := rotatelogs.WithRotationCount(l.cfg.LogRotationCount)
		if l.cfg.LogRotationTime == 0 {
			l.cfg.LogRotationTime = 60 * 60 * 24
		}
		optWithRotationTime := rotatelogs.WithRotationTime(time.Duration(l.cfg.LogRotationTime) * time.Second)
		var opts []rotatelogs.Option
		if runtime.GOOS != constant.GOOSWindows {
			opts = append(opts, optWithLinkName)
		}
		opts = append(opts, optWithMaxAge)
		if l.cfg.LogRotationCount > 0 {
			opts = append(opts, optWithRotationCount)
		} else {
			opts = append(opts, optWithRotationTime)
		}
		fd, _ = rotatelogs.New(optLogFileFmt, opts...)

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
