package util

import (
	"log/syslog"

	prefixed "git.aventer.biz/AVENTER/go-logrus-formatter"
	"github.com/sirupsen/logrus"
	logrusSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

// SetLogging sets the loglevel and text formating
func SetLogging(level logrus.Level, enableSyslog bool, name string) {
	logrus.SetFormatter(&prefixed.TextFormatter{
		ForceColors: true,
	})
	logrus.SetLevel(level)

	if enableSyslog {
		hook, _ := logrusSyslog.NewSyslogHook("", "", syslog.LOG_DEBUG, name)
		logrus.AddHook(hook)
	}
}
