package util

import (
	"log/syslog"

	prefixed "git.aventer.biz/AVENTER/go-logrus-formatter"
	Log "github.com/sirupsen/logrus"
	logrusSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

// SetLogging sets the loglevel and text formating
func SetLogging(level Log.Level, enableSyslog bool, name string) {
	Log.SetFormatter(&prefixed.TextFormatter{
		ForceColors: true,
	})
	Log.SetLevel(level)

	if enableSyslog {
		hook, _ := logrusSyslog.NewSyslogHook("", "", syslog.LOG_DEBUG, name)
		Log.AddHook(hook)
	}
}
