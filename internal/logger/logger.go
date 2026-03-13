package logger

import (
	"fmt"
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
	s := fmt.Sprintf(msg+"\n", args...)
	sl.infoLog.Output(2, s)
}

func (sl *StdLogger) Warn(msg string, args ...interface{}) {
	s := fmt.Sprintf(msg+"\n", args...)
	sl.warnLog.Output(2, s)
}

func (sl *StdLogger) Error(msg string, args ...interface{}) {
	s := fmt.Sprintf(msg+"\n", args...)
	sl.errorLog.Output(2, s)
}

func (sl *StdLogger) Fatal(msg string, args ...interface{}) {
	s := fmt.Sprintf(msg+"\n", args...)
	sl.fatalLog.Output(2, s)
}

func (sl *StdLogger) Debug(msg string, args ...interface{}) {
	if sl.debugMode {
		s := fmt.Sprintf(msg+"\n", args...)
		sl.debugLog.Output(2, s)
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
