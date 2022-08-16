package logger

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelInfo LogLevel = iota
	LogLevelWarning
	LogLevelError
)

type Logger struct {
	logger *log.Logger
	level  LogLevel
}

func NewLogExtended() *Logger {
	return &Logger{
		logger: log.New(os.Stderr, "", log.LstdFlags),
		level:  LogLevelInfo,
	}
}

func (l *Logger) SetLogLevel(logLevel LogLevel) {
	l.level = logLevel
}

func (l *Logger) Infoln(msg string) {
	l.println(0, "INFO", msg)
}

func (l *Logger) Warnln(msg string) {
	l.println(1, "WARN", msg)
}

func (l *Logger) Errorln(msg string) {
	l.println(2, "ERR", msg)
}

func (l *Logger) Println(msg string) {
	l.logger.Println(msg)
}

func (l *Logger) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.level > srcLogLvl {
		return
	}

	l.logger.Println(prefix, msg)
}
