package log

import "github.com/sirupsen/logrus"

type logrusLogger struct {
	trace *logrusSubLog
	debug *logrusSubLog
	info  *logrusSubLog
	warn  *logrusSubLog
	error *logrusSubLog
}

type logrusSubLog struct {
	level       Level
	logrusLevel logrus.Level
}

func NewLogrusLogger() *logrusLogger {
	return &logrusLogger{
		trace: &logrusSubLog{
			level:       TraceLevel,
			logrusLevel: logrus.TraceLevel,
		},
		debug: &logrusSubLog{
			level:       DebugLevel,
			logrusLevel: logrus.DebugLevel,
		},
		info: &logrusSubLog{
			level:       InfoLevel,
			logrusLevel: logrus.InfoLevel,
		},
		warn: &logrusSubLog{
			level:       WarnLevel,
			logrusLevel: logrus.WarnLevel,
		},
		error: &logrusSubLog{
			level:       ErrorLevel,
			logrusLevel: logrus.ErrorLevel,
		},
	}
}

func (l *logrusLogger) Trace() SubLog {
	return l.trace
}

func (l *logrusLogger) Debug() SubLog {
	return l.trace
}

func (l *logrusLogger) Info() SubLog {
	return l.trace
}

func (l *logrusLogger) Warn() SubLog {
	return l.trace
}

func (l *logrusLogger) Error() SubLog {
	return l.trace
}

func (s *logrusSubLog) Level() Level {
	return s.level
}

func (s *logrusSubLog) IsEnable() bool {
	return logrus.IsLevelEnabled(s.logrusLevel)
}
func (s *logrusSubLog) Log(args ...interface{}) {
	switch s.level {
	case TraceLevel:
		logrus.Trace(args)
	case DebugLevel:
		logrus.Debug(args)
	case InfoLevel:
		logrus.Info(args)
	case WarnLevel:
		logrus.Warn(args)
	case ErrorLevel:
		logrus.Error(args)
	}
}
func (s *logrusSubLog) Logf(format string, args ...interface{}) {
	switch s.level {
	case TraceLevel:
		logrus.Tracef(format, args)
	case DebugLevel:
		logrus.Debugf(format, args)
	case InfoLevel:
		logrus.Infof(format, args)
	case WarnLevel:
		logrus.Warnf(format, args)
	case ErrorLevel:
		logrus.Errorf(format, args)
	}
}
func (s *logrusSubLog) Logln(args ...interface{}) {
	switch s.level {
	case TraceLevel:
		logrus.Traceln(args)
	case DebugLevel:
		logrus.Debugln(args)
	case InfoLevel:
		logrus.Infoln(args)
	case WarnLevel:
		logrus.Warnln(args)
	case ErrorLevel:
		logrus.Errorln(args)
	}
}
