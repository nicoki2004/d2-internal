package logger

import (
	"log"
	"os"
)

// Logger abstrae la interfaz de logging
type Logger interface {
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
}

// StdLogger implements Logger using standard log
type StdLogger struct {
	infoLog   *log.Logger
	warnLog   *log.Logger
	errorLog  *log.Logger
	fatalLog  *log.Logger
	debugLog  *log.Logger
	debugMode bool
}

// NewStdLogger creates a new standard logger
func NewStdLogger(debugMode bool) *StdLogger {
	return &StdLogger{
		infoLog:   log.New(os.Stdout, "✓ ", log.Lshortfile),
		warnLog:   log.New(os.Stdout, "⚠ ", log.Lshortfile),
		errorLog:  log.New(os.Stderr, "✗ ", log.Lshortfile),
		fatalLog:  log.New(os.Stderr, "💥 ", log.Lshortfile),
		debugLog:  log.New(os.Stdout, "🐛 ", log.Lshortfile),
		debugMode: debugMode,
	}
}

func (sl *StdLogger) Info(msg string, args ...interface{}) {
	sl.infoLog.Printf(msg, args...)
}

func (sl *StdLogger) Warn(msg string, args ...interface{}) {
	sl.warnLog.Printf(msg, args...)
}

func (sl *StdLogger) Error(msg string, args ...interface{}) {
	sl.errorLog.Printf(msg, args...)
}

func (sl *StdLogger) Fatal(msg string, args ...interface{}) {
	sl.fatalLog.Fatalf(msg, args...)
}

func (sl *StdLogger) Debug(msg string, args ...interface{}) {
	if sl.debugMode {
		sl.debugLog.Printf(msg, args...)
	}
}

// Global logger instance
var globalLogger Logger = NewStdLogger(false)

func GetLogger() Logger {
	return globalLogger
}

func SetLogger(l Logger) {
	globalLogger = l
}
