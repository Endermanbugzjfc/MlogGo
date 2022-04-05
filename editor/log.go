package editor

import "github.com/sirupsen/logrus"

func InitLogrus(logger *logrus.Logger) {
	formatter := logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,

		FullTimestamp: true,

		TimestampFormat:        "",
		SortingFunc:            nil,
		DisableSorting:         false,
		DisableLevelTruncation: false,
		PadLevelText:           false,
		QuoteEmptyFields:       false,
		FieldMap:               logrus.FieldMap{},
		CallerPrettyfier:       nil,
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

	SetDebugMode(enabled bool)
}

func LogrusToEditorLogger(logger *logrus.Logger) Logger {
	return logrusWrapper{logger}
}

type logrusWrapper struct {
	*logrus.Logger
}

func (wrapper logrusWrapper) SetDebugMode(enabled bool) {
	log := wrapper.Logger
	if enabled {
		log.SetLevel(logrus.TraceLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
}
