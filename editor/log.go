package editor

import "github.com/sirupsen/logrus"

func InitLogrus(logger logrus.Logger) {
	formatter := logrus.TextFormatter{
		FullTimestamp: true,
	}
	logger.SetFormatter(&formatter)
}

type Logger interface {
	Tracef(format string, args ...any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Panicf(format string, args ...any)
}
