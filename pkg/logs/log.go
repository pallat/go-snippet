package logs

import (
	"fmt"
	"io"
)

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
	return Logger{w: w}
}

type Logger struct {
	w io.Writer
}

func (Logger) Info(args ...interface{}) {}
func (l Logger) Debug(args ...interface{}) {
	fmt.Fprintln(l.w, args...)
}
func (Logger) Error(args ...interface{}) {}
