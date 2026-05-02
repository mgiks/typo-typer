package logger

import "log/slog"

type LoggerService struct {
	logger slog.Logger
}

func NewService(logger slog.Logger) LoggerService {
	return LoggerService{
		logger: logger,
	}
}

func (s LoggerService) Error(msg string, args ...any) {
	s.logger.Error(msg, args...)
}
