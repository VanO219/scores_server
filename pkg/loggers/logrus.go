package loggers

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logrus struct {
	lgInfo  *logrus.Logger
	lgError *logrus.Logger
}

func (l *Logrus) Error(input ...interface{}) {
	l.lgError.Errorln(input)
}

func (l *Logrus) Errorf(format string, args ...interface{}) {
	l.lgError.Errorf(format, args)
}

func (l *Logrus) Info(input ...interface{}) {
	l.lgInfo.Infoln(input)
}

func (l *Logrus) Infof(format string, args ...interface{}) {
	l.lgError.Infof(format, args)
}

func (l *Logrus) Debug(input ...interface{}) {
	l.lgError.Debugln(input)
}

func (l *Logrus) Debugf(format string, args ...interface{}) {
	l.lgError.Debugf(format, args)
}

func NewLogrus(outputErrorsFile *os.File, outputInfosFile *os.File, formatter logrus.Formatter) *Logrus {
	le := logrus.New()
	le.SetOutput(outputErrorsFile)
	le.SetFormatter(formatter)

	li := logrus.New()
	li.SetOutput(outputInfosFile)
	li.SetFormatter(formatter)

	lg := &Logrus{
		lgInfo:  li,
		lgError: le,
	}
	return lg
}
