package logging

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggerFactory struct {
	*logrus.Logger
	serviceCode string
}

const (
	SERVICE_CODE string = "SERVICECODE"

	INFO  uint32 = uint32(logrus.InfoLevel)
	DEBUG uint32 = uint32(logrus.DebugLevel)
	ERROR uint32 = uint32(logrus.ErrorLevel)
	PANIC uint32 = uint32(logrus.PanicLevel)
	WARN  uint32 = uint32(logrus.WarnLevel)
	FATAL uint32 = uint32(logrus.FatalLevel)
	TRACE uint32 = uint32(logrus.TraceLevel)
)

// Default or root logger
var defaultLoggerFactory *LoggerFactory

// serviceCode used to identify logs of a service
// when logs are managed in some central logging
// service like ELK.
func Init(serviceCode string, output io.Writer, logLevel uint32) {
	defaultLoggerFactory = new(LoggerFactory)
	defaultLoggerFactory.Logger = logrus.New()
	defaultLoggerFactory.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	}

	if output == nil {
		output = os.Stdout
	}

	defaultLoggerFactory.SetOutput(output)
	defaultLoggerFactory.serviceCode = serviceCode
	defaultLoggerFactory.Level = logrus.Level(logLevel)
}

// Logger Interface
type LoggerInterface interface {
	// Infof log at info level
	Infof(fmt string, args ...interface{})

	// Warnf log at warn level
	Warnf(fmt string, args ...interface{})

	// Panicf log at Panic level
	Panicf(fmt string, args ...interface{})

	// Tracef log at Trace level
	Tracef(fmt string, args ...interface{})

	// Fatalf log at Fatal level
	Fatalf(fmt string, args ...interface{})

	// Debugf log at Debug level
	Debugf(fmt string, args ...interface{})

	// Errorf log at Debug level
	Errorf(fmt string, args ...interface{})

	// Eventf log with event marker
	Eventf(verb, subject, object string, level uint32, fmt string, args ...interface{})

	// Notifyf log with notify marker
	Notifyf(verb, subject, object string, level uint32, fmt string, args ...interface{})
}

type Logger struct {
	*logrus.Entry
}

// NewLogger returns a new logger with given log level.
func NewLogger(logLevel uint32) *Logger {
	logger := logrus.NewEntry(defaultLoggerFactory.Logger).WithField(
		SERVICE_CODE, defaultLoggerFactory.serviceCode)
	return &Logger{
		Entry: logger,
	}
}

// Infof log at info level
func (lg *Logger) Infof(fmt string, args ...interface{}) {
	lg.Entry.WithField("TYPE", "LOG").Infof(fmt, args...)
}

// Info log at info level
func (lg *Logger) Info(fmt string) {
	lg.Entry.WithField("TYPE", "LOG").Info(fmt)
}

// Infoln log at info level
func (lg *Logger) Infoln(fmt string) {
	lg.Entry.WithField("TYPE", "LOG").Infoln(fmt + "\n")
}

// Warnf log at warn level
func (lg *Logger) Warnf(fmt string, args ...interface{}) {
	lg.Entry.WithField("TYPE", "LOG").Warnf(fmt, args...)
}

// Warn log at warn level
func (lg *Logger) Warn(fmt string) {
	lg.Entry.WithField("TYPE", "LOG").Warn(fmt)
}

// Warnln log at warn level
func (lg *Logger) Warnln(fmt string) {
	lg.Entry.WithField("TYPE", "LOG").Warnln(fmt + "\n")
}

// Panicf log at Panic level
func (lg *Logger) Panicf(fmt string, args ...interface{}) {
	lg.Entry.WithField("TYPE", "LOG").Panicf(fmt, args...)
}

// Tracef log at Trace level
func (lg *Logger) Tracef(fmt string, args ...interface{}) {
	lg.Entry.WithField("TYPE", "LOG").Tracef(fmt, args...)
}

// Fatalf log at Fatal level
func (lg *Logger) Fatalf(fmt string, args ...interface{}) {
	lg.Entry.WithField("TYPE", "LOG").Fatalf(fmt, args...)
}

// Debugf log at Debug level
func (lg *Logger) Debugf(fmt string, args ...interface{}) {
	lg.Entry.WithField("TYPE", "LOG").Debugf(fmt, args...)
}

// Debug log at debug level
func (lg *Logger) Debug(fmt string) {
	lg.Entry.WithField("TYPE", "LOG").Debug(fmt)
}

// Debugln log at warn level
func (lg *Logger) Debugln(fmt string) {
	lg.Entry.WithField("TYPE", "LOG").Debugln(fmt + "\n")
}

// Errorf log at Debug level
func (lg *Logger) Errorf(fmt string, args ...interface{}) {
	lg.Entry.WithField("TYPE", "LOG").Errorf(fmt, args...)
}

// Eventf log with event marker
func (lg *Logger) Eventf(verb, subject, object string,
	level uint32, fmt string, args ...interface{}) {

	entry := lg.Entry.WithField("TYPE", "EVENT").
		WithField("VERB", verb).
		WithField("SUBJECT", subject).
		WithField("OBJECT", object)
	switch level {
	case INFO:
		entry.Infof(fmt, args...)
	case DEBUG:
		entry.Debugf(fmt, args...)
	case ERROR:
		entry.Errorf(fmt, args...)
	case TRACE:
		entry.Tracef(fmt, args...)
	case PANIC:
		entry.Panicf(fmt, args...)
	case FATAL:
		entry.Fatalf(fmt, args...)
	case WARN:
		entry.Warnf(fmt, args...)
	}
}

// Notifyf log with notify marker
func (lg *Logger) Notifyf(verb, subject, object string, level uint32,
	fmt string, args ...interface{}) {

	entry := lg.Entry.WithField("TYPE", "NOTIFY").
		WithField("VERB", verb).
		WithField("SUBJECT", subject).
		WithField("OBJECT", object)

	switch level {
	case INFO:
		entry.Infof(fmt, args...)
	case DEBUG:
		entry.Debugf(fmt, args...)
	case ERROR:
		entry.Errorf(fmt, args...)
	case TRACE:
		entry.Tracef(fmt, args...)
	case PANIC:
		entry.Panicf(fmt, args...)
	case FATAL:
		entry.Fatalf(fmt, args...)
	case WARN:
		entry.Warnf(fmt, args...)
	}
}

// Register a log hook
func RegisterLogHook(hook logrus.Hook) {
	defaultLoggerFactory.Hooks.Add(hook)
}

// Make sure that Logger implements all func of logger interface
var _ LoggerInterface = &Logger{}
