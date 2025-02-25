package logrus

import (
	"url-shortener/pkg/log"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogger(logLevel string) log.Logger {
	log := logrus.New()

	if logLevel == "" {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Invalid log level: %s", logLevel)
	}
	log.SetLevel(level)

	return &LogrusLogger{logger: log}
}

func (l *LogrusLogger) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *LogrusLogger) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *LogrusLogger) Fatal(args ...any) {
	l.logger.Fatal(args...)
}
