package logs

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

func New() Logger {
	return Logger{}
}

type Logger struct{}

func (Logger) Info(args ...interface{})  {}
func (Logger) Debug(args ...interface{}) {}
func (Logger) Error(args ...interface{}) {}
