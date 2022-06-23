package logging

import (
	"fmt"

	glogger "cloud.google.com/go/logging"
)

type Log struct {
	Logger *glogger.Logger
}

func (l *Log) Info(m string) {
	lg := l.Logger.StandardLogger(glogger.Info)
	lg.Println(m)
}

func (l *Log) Error(m string, e error) {
	lg := l.Logger.StandardLogger(glogger.Error)
	o := fmt.Sprintf("%s: %e", m, e)
	lg.Println(o)
}

func (l *Log) Critical(m string, e error) {
	lg := l.Logger.StandardLogger(glogger.Critical)

	if e != nil {
		o := fmt.Sprintf("%s: %e", m, e)
		lg.Println(o)
		return
	}

	lg.Println(m)
}
