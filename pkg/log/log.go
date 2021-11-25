package log

import (
	"log"

	"go.uber.org/zap"
)

type Field = zap.Field

var defaultLogger zap.Logger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defaultLogger = *logger
}

func Info(msg string, fields ...Field) {
	defaultLogger.Info(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	defaultLogger.Fatal(msg, fields...)
}

func Error(msg string, fields ...Field) {
	defaultLogger.Error(msg, fields...)
}

func ErrorField(err error) Field {
	return zap.Error(err)
}
