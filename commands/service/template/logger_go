package {{.Name}}

import (
	"github.com/sirupsen/logrus"
	"github.com/evalphobia/logrus_sentry"
	"os"
	"time"
)

var (
	// Log global settings of logger
	Log *logrus.Logger
	dsn = "{{.Dsn}}"
)

func init() {

	Log = logrus.New()

	// Logging format is JSON
	Log.Formatter = &logrus.JSONFormatter{}

	Log.Out = os.Stdout

	levels := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
	hook, err := logrus_sentry.NewSentryHook(dsn, levels)
	hook.Timeout = 20 * time.Second
	hook.StacktraceConfiguration.Enable = true

	if err == nil {
		Log.Hooks.Add(hook)
	}
}
