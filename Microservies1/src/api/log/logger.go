package log

import (
	"os"
	"src/api/config"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}
	if config.IsProd() {
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		Log.Formatter = &logrus.TextFormatter{}

	}
}

func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Info(msg)

}
func Error(msg string, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Error(msg)

}

func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Debug(msg)

}


func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		ele := strings.Split(tag, ":")
		result[strings.TrimSpace(ele[0])] = strings.TrimSpace(ele[1])
	}
	return result
}
