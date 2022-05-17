package log

import "github.com/sirupsen/logrus"

// Log log
// @todo waiting to complete
type Log struct{}

func (l Log) Print(args ...interface{}) {
	logrus.Print(args)
}

func (l Log) Info(args ...interface{}) {
	logrus.Info(args)
}

func (l Log) Error(args ...interface{}) {
	logrus.Error(args)
}
