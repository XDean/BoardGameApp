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
	logger      *logrus.Logger
}

func NewLogrusLogger(logger *logrus.Logger) *logrusLogger {
	return &logrusLogger{
		trace: &logrusSubLog{
			level:       TraceLevel,
			logrusLevel: logrus.TraceLevel,
			logger:      logger,
		},
		debug: &logrusSubLog{
			level:       DebugLevel,
			logrusLevel: logrus.DebugLevel,
			logger:      logger,
		},
		info: &logrusSubLog{
			level:       InfoLevel,
			logrusLevel: logrus.InfoLevel,
			logger:      logger,
		},
		warn: &logrusSubLog{
			level:       WarnLevel,
			logrusLevel: logrus.WarnLevel,
			logger:      logger,
		},
		error: &logrusSubLog{
			level:       ErrorLevel,
			logrusLevel: logrus.ErrorLevel,
			logger:      logger,
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
		s.logger.Trace(args)
	case DebugLevel:
		s.logger.Debug(args)
	case InfoLevel:
		s.logger.Info(args)
	case WarnLevel:
		s.logger.Warn(args)
	case ErrorLevel:
		s.logger.Error(args)
	}
}
func (s *logrusSubLog) Logf(format string, args ...interface{}) {
	switch s.level {
	case TraceLevel:
		s.logger.Tracef(format, args)
	case DebugLevel:
		s.logger.Debugf(format, args)
	case InfoLevel:
		s.logger.Infof(format, args)
	case WarnLevel:
		s.logger.Warnf(format, args)
	case ErrorLevel:
		s.logger.Errorf(format, args)
	}
}
func (s *logrusSubLog) Logln(args ...interface{}) {
	switch s.level {
	case TraceLevel:
		s.logger.Traceln(args)
	case DebugLevel:
		s.logger.Debugln(args)
	case InfoLevel:
		s.logger.Infoln(args)
	case WarnLevel:
		s.logger.Warnln(args)
	case ErrorLevel:
		s.logger.Errorln(args)
	}
}
