package logger

type Loggers interface {
	Error(input ...interface{})
	Errorf(format string, args ...interface{})
	Info(input ...interface{})
	Infof(format string, args ...interface{})
	Debug(input ...interface{})
	Debugf(format string, args ...interface{})
}
