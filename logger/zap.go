package logger

import (
	"log"

	"go.uber.org/zap"
)

func NewZap() *zap.Logger {
	config := zap.NewProductionConfig()
	logger, err := config.Build(zap.AddCaller())
	if err != nil {
		log.Fatal(err)
	}

	return logger
}
