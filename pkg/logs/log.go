package logs

import "io"

type Level int

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func New(w io.Writer) Logger {
	return Logger{}
}

type Logger struct{}

func (Logger) Info(args ...interface{})  {}
func (Logger) Debug(args ...interface{}) {}
func (Logger) Error(args ...interface{}) {}
