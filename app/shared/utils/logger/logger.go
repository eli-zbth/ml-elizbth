package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Entry
}

type Fields logrus.Fields

func InitLoggerConfig() {
	setLevel()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func setLevel() {
	//logLevel, err := logrus.ParseLevel(viper.GetString("LOG_LEVEL"))
	logLevel, err := logrus.ParseLevel("INFO")
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)
}

func New() Logger {
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

func (l *Logger) Infoln(args ...interface{}) {
	l.log.Infoln(args...)
}
