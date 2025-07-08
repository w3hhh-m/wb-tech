package logger

const (
	ErrorKey = "error"
)

// Logger interface
type Logger interface {
	Debug(msg string, fields ...LogField)
	Info(msg string, fields ...LogField)
	Warn(msg string, fields ...LogField)
	Error(msg string, fields ...LogField)
	Fatal(msg string, fields ...LogField)
	Panic(msg string, fields ...LogField)

	With(fields ...LogField) Logger
	Sync() error
}

// LogField is used in Logger to contextual logging
type LogField struct {
	Key   string
	Value interface{}
}

// Field creates and returns LogField used in Logger
func Field(key string, value interface{}) LogField {
	return LogField{Key: key, Value: value}
}

// Error creates and returns LogField with ErrorKey used in Logger
func Error(value interface{}) LogField {
	return LogField{Key: ErrorKey, Value: value}
}
