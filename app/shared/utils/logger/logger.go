package logger

import (
	"ml-elizabeth/app/domain/entity"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Entry
	Trace *entity.Trace
}

type Fields logrus.Fields

func InitLoggerConfig() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func NewLogger() Logger {
	logrus.SetLevel(logrus.DebugLevel)
	l := Logger{
		log: logrus.NewEntry(logrus.StandardLogger()),
	}
	return l
}

func (l *Logger) WithFields(fields Fields) {
	l.log = l.log.WithFields(logrus.Fields(fields))
}

func (l *Logger) Info(args ...interface{}) {
	l.log.Info(args...)
}


func (l *Logger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}


func (l *Logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func NewWithTrace(trace *entity.Trace) Logger {
	l := Logger{
		log:   logrus.NewEntry(logrus.StandardLogger()),
		Trace: trace,
	}

	if trace != nil {
		l.WithFields(Fields{
			"traceId":  trace.TraceID,
			"spanId":   trace.SpanID,
			"parentId": trace.ParentID,
		})
	}

	return l
}
