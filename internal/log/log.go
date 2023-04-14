package loggeri2

import (
	"go.uber.org/zap"
)

func NewLog() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
