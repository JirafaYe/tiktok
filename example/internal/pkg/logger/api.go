package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Info(string, ...zap.Field)
	Fatal(string, ...zap.Field)
	Debug(string, ...zap.Field)
	Error(string, ...zap.Field)
}

func (m *Manager) SyncLogger() {
	m.logger.Sync()
}

// Info level log
func (m *Manager) Info(msg string, fields ...zap.Field) {
	m.logger.Info(msg, fields...)
}

// Fatal level log
func (m *Manager) Fatal(msg string, fields ...zap.Field) {
	m.logger.Fatal(msg, fields...)
}

// Debug level log
func (m *Manager) Debug(msg string, fields ...zap.Field) {
	m.logger.Debug(msg, fields...)
}

// Error level log
func (m *Manager) Error(msg string, fields ...zap.Field) {
	m.logger.Error(msg, fields...)
}

func (m *Manager) String(cause string) zap.Field {
	return zap.String("cause", cause)
}
