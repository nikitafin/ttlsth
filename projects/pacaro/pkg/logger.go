package logger

import "log"

type Logger interface {
	Init() error
	Debug(args ...any)
	Fatal(args ...any)
}

type dummyLogger struct {
	l log.Logger
}

// Init implements Logger
func (*dummyLogger) Init() error {
	// TODO(Nikita): add init
	return nil
}

// Debug implements Logger
func (r *dummyLogger) Debug(args ...any) {
	r.l.Println()
}

// Fatal implements Logger
func (r *dummyLogger) Fatal(args ...any) {
	panic("unimplemented")
}

func NewLogger() Logger {
	return &dummyLogger{
		l: *log.Default(),
	}
}
