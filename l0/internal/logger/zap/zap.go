package zaplogger

import (
	"fmt"
	"wb-tech-l0/internal/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap is a Logger interface implementation for Zap
type Zap struct {
	logger *zap.Logger
}

// New creates and returns initialized Zap implementation of Logger interface
func New(level, hostname string) (*Zap, error) {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		return nil, fmt.Errorf("failed to parse log level: %w", err)
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	cfg := zap.NewProductionConfig()

	encoderCfg.TimeKey = "timestamp"
	encoderCfg.LevelKey = "level"
	encoderCfg.NameKey = "logger"
	encoderCfg.CallerKey = "caller"
	encoderCfg.MessageKey = "message"
	encoderCfg.StacktraceKey = "stacktrace"
	encoderCfg.LineEnding = zapcore.DefaultLineEnding
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeDuration = zapcore.StringDurationEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder

	cfg.Level = zap.NewAtomicLevelAt(zapLevel)
	cfg.EncoderConfig = encoderCfg
	cfg.Sampling = nil
	cfg.InitialFields = map[string]interface{}{
		"hostname": hostname,
	}

	options := []zap.Option{
		zap.AddCallerSkip(1),
	}

	log, err := cfg.Build(options...)
	if err != nil {
		return nil, fmt.Errorf("failed to build zap logger: %w", err)
	}
	return &Zap{logger: log}, nil
}

func (l *Zap) Debug(msg string, fields ...logger.LogField) {
	l.logger.Debug(msg, convertFields(fields)...)
}

func (l *Zap) Info(msg string, fields ...logger.LogField) {
	l.logger.Info(msg, convertFields(fields)...)
}

func (l *Zap) Warn(msg string, fields ...logger.LogField) {
	l.logger.Warn(msg, convertFields(fields)...)
}

func (l *Zap) Error(msg string, fields ...logger.LogField) {
	l.logger.Error(msg, convertFields(fields)...)
}

func (l *Zap) Fatal(msg string, fields ...logger.LogField) {
	l.logger.Fatal(msg, convertFields(fields)...)
}

func (l *Zap) Panic(msg string, fields ...logger.LogField) {
	l.logger.Panic(msg, convertFields(fields)...)
}

func (l *Zap) With(fields ...logger.LogField) logger.Logger {
	return &Zap{logger: l.logger.With(convertFields(fields)...)}
}

func (l *Zap) Sync() error {
	return l.logger.Sync()
}

func convertFields(fields []logger.LogField) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}
