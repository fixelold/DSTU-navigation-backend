package logging

import "github.com/sirupsen/logrus"

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func init() {
	l := logrus.New()
	e = logrus.NewEntry(l)
}
