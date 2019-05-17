package log

import (
	"errors"
	"strings"
)

var Log Logger

type Level uint32

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
)

type Logger interface {
	Trace() SubLog
	Debug() SubLog
	Info() SubLog
	Warn() SubLog
	Error() SubLog
}

type SubLog interface {
	Level() Level
	IsEnable() bool
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Logln(args ...interface{})
}

func ParseLevel(level string) (Level, error) {
	switch strings.ToLower(level) {
	case "trace":
		return TraceLevel, nil
	case "debug":
		return DebugLevel, nil
	case "info":
		return InfoLevel, nil
	case "warn":
		return WarnLevel, nil
	case "error":
		return ErrorLevel, nil
	default:
		return InfoLevel, errors.New("Unsupported Log level: " + level)
	}
}

func (level Level) String() string {
	switch level {
	case TraceLevel:
		return "TRACE"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARNING"
	case ErrorLevel:
		return "ERROR"
	}
	return "UNKNOWN"
}

func Trace() SubLog {
	return Log.Trace()
}

func Debug() SubLog {
	return Log.Debug()
}

func Info() SubLog {
	return Log.Info()
}

func Warn() SubLog {
	return Log.Warn()
}

func Error() SubLog {
	return Log.Error()
}
