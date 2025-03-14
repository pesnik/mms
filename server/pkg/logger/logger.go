package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(format string, v ...interface{})
	Error(format string, v ...interface{})
	Debug(format string, v ...interface{})
}

type SimpleLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

var Log Logger

func New() Logger {
	return &SimpleLogger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *SimpleLogger) Info(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

func (l *SimpleLogger) Error(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

func (l *SimpleLogger) Debug(format string, v ...interface{}) {
	l.debugLogger.Printf(format, v...)
}
