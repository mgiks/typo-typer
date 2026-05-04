package logger

import (
	"log/slog"
	"os"
)

type LoggerService interface {
	Info(string, ...any)
	Error(string, ...any)
	FatalError(string, ...any)
}

type loggerService struct {
	logger slog.Logger
}

func NewService(logger slog.Logger) LoggerService {
	return loggerService{
		logger: logger,
	}
}

func (s loggerService) Info(msg string, args ...any) {
	s.logger.Info(msg, args...)
}

func (s loggerService) Error(msg string, args ...any) {
	s.logger.Error(msg, args...)
}

func (s loggerService) FatalError(msg string, args ...any) {
	s.logger.Error(msg, args...)
	os.Exit(1)
}
