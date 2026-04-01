package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)


func NewLogger() *logrus.Logger{
	log := logrus.New()
	
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
	
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	return log
}