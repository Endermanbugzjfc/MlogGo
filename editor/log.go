package editor

import "github.com/sirupsen/logrus"

func InitLogger(logger logrus.Logger) {
	formatter := logrus.TextFormatter{
		FullTimestamp: true,
	}
	logger.SetFormatter(&formatter)
}
