package logging

import (
	"context"
	"fmt"
	"log"

	glogger "cloud.google.com/go/logging"
)

type Log struct {
	logger *glogger.Logger
}

func (l *Log) NewLogger(projectID string, logName string) {
	client, err := glogger.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	l.logger = client.Logger(logName)
}

func (l *Log) Info(m string) {
	lg := l.logger.StandardLogger(glogger.Info)
	lg.Println(m)
}

func (l *Log) Error(m string, e error) {
	lg := l.logger.StandardLogger(glogger.Error)
	o := fmt.Sprintf("%s: %e", m, e)
	lg.Println(o)
}

func (l *Log) Critical(m string, e error) {
	lg := l.logger.StandardLogger(glogger.Critical)

	if e != nil {
		o := fmt.Sprintf("%s: %e", m, e)
		lg.Println(o)
		return
	}

	lg.Println(m)
}
