package editor

import (
	"flag"

	"github.com/sirupsen/logrus"
)

// Init runs flag.Parse().
// So please register all your own flags before calling.
func Init() {
	logger = LogrusToEditorLogger(logrus.StandardLogger())

	configArgument := RegisterConfigArgument()
	flag.Parse()

	MustLoadConfig(logger, *configArgument)
}
