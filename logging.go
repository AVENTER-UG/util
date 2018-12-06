package util

import (
	"log/syslog"

	prefixed "git.aventer.biz/AVENTER/go-logrus-formatter"
	log "github.com/sirupsen/logrus"
	logrusSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

// SetLogging sets the loglevel and text formating
func SetLogging(level log.Level, enableSyslog bool, name string) {
	log.SetFormatter(&prefixed.TextFormatter{
		ForceColors: true,
	})
	log.SetLevel(level)

	if enableSyslog {
		hook, _ := logrusSyslog.NewSyslogHook("", "", syslog.LOG_DEBUG, name)
		log.AddHook(hook)
	}
}
