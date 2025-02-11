package log

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(logLevel string) *logrus.Logger {
	log := logrus.New()

	if logLevel == "" {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Invalid log level: %s", logLevel)
	}
	log.SetLevel(level)

	return log
}
