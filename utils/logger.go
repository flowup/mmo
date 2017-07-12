package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/evalphobia/logrus_sentry"
	"os"
	"time"
)

var (
	// Log global settings of logger
	Log *logrus.Logger
	dsn = "https://5554d201ce064e8790792540de39c608:fb93a01a320a486ab40b4fbb5feaf7ac@sentry.io/190135"
)

func init() {

	Log = logrus.New()

	// Logging format is Text
	Log.Formatter = &logrus.TextFormatter{}

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
